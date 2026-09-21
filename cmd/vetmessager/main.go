package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	"github.com/maximyunak/vetmessager/internal/core/password"
	core_postgres_pool "github.com/maximyunak/vetmessager/internal/core/repository/postgres/pull"
	core_http_middleware "github.com/maximyunak/vetmessager/internal/core/transport/http/middleware"
	core_http_server "github.com/maximyunak/vetmessager/internal/core/transport/http/server"
	users_postgres_repository "github.com/maximyunak/vetmessager/internal/features/users/repository/postgres"
	users_service "github.com/maximyunak/vetmessager/internal/features/users/service"
	users_transport_http "github.com/maximyunak/vetmessager/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// logger
	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()

	// connection pool
	logger.Debug("Initializing postgtes connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(ctx, core_postgres_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("Failed to connect to postgres", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	passwordHasher := password.NewBcryptHasher(password.NewConfigMust())
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository, passwordHasher)

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)
	userRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(userRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterApiRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("failed to start http server", zap.Error(err))
	}
}
