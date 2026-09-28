package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/maximyunak/vetmessager/internal/core/auth"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	"github.com/maximyunak/vetmessager/internal/core/password"
	core_postgres_pool "github.com/maximyunak/vetmessager/internal/core/repository/postgres/pull"
	core_http_middleware "github.com/maximyunak/vetmessager/internal/core/transport/http/middleware"
	core_http_server "github.com/maximyunak/vetmessager/internal/core/transport/http/server"
	chat_postgres_repository "github.com/maximyunak/vetmessager/internal/features/chat/repository/postgres"
	messages_postgres_repository "github.com/maximyunak/vetmessager/internal/features/messages/repository/postgres"
	messages_service "github.com/maximyunak/vetmessager/internal/features/messages/service"
	messages_transport_http "github.com/maximyunak/vetmessager/internal/features/messages/transport/http"
	users_postgres_repository "github.com/maximyunak/vetmessager/internal/features/users/repository/postgres"
	users_service "github.com/maximyunak/vetmessager/internal/features/users/service"
	users_transport_http "github.com/maximyunak/vetmessager/internal/features/users/transport/http"
	"github.com/maximyunak/vetmessager/internal/realtime/websocket"
	"go.uber.org/zap"

	_ "github.com/maximyunak/vetmessager/docs"
)

// @title           VetMessager
// @version         1.0
// @description    	Sveta the best
// @host      localhost:5050
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your JWT token with the Bearer prefix. Example: Bearer {token}
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

	// jwt
	tokenManager := auth.NewJWTManager(auth.NewConfigMust())
	checkAuth := core_http_middleware.CheckAuth(tokenManager)

	// user feature
	logger.Debug("initializing feature", zap.String("feature", "users"))
	passwordHasher := password.NewBcryptHasher(password.NewConfigMust())
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository, passwordHasher, tokenManager)

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)
	userPublicRoutes := usersTransportHTTP.PublicRoutes()
	userProtectedRoutes := usersTransportHTTP.ProtectedRoutes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(userPublicRoutes...)
	apiVersionRouter.RegisterProtectedRoutes(checkAuth, userProtectedRoutes...)

	// chat feature

	chatRepository := chat_postgres_repository.NewChatRepository(pool)

	// websocket conn

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub, chatRepository)
	go hub.Run()

	messageWsProtectedRoutes := wsHandler.ProtectedRoutes()

	apiVersionRouter.RegisterProtectedRoutes(
		checkAuth,
		messageWsProtectedRoutes...)

	// messages feature
	logger.Debug("initializing feature", zap.String("feature", "messages"))

	messagesRepository := messages_postgres_repository.NewMessagesRepository(pool)
	messagesService := messages_service.NewMessageService(messagesRepository, tokenManager)
	messagesTransportHTTP := messages_transport_http.NewMessagesHTTPHandler(messagesService)

	messageProtectedRoutes := messagesTransportHTTP.ProtectedRoutes()
	apiVersionRouter.RegisterProtectedRoutes(checkAuth, messageProtectedRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterApiRouters(apiVersionRouter)
	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("failed to start http server", zap.Error(err))
	}
}
