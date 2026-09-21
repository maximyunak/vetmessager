package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/maximyunak/vetmessager/docs"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_middleware "github.com/maximyunak/vetmessager/internal/core/transport/http/middleware"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	log    *core_logger.Logger

	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, middleware ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (h *HTTPServer) RegisterApiRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		pref := "/api/" + string(router.apiVersion)

		h.mux.Handle(pref+"/", http.StripPrefix(pref, router))
	}
}

func (h *HTTPServer) RegisterSwagger() {
	h.mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	h.mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
	})
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)
	go func() {
		defer close(ch)

		h.log.Warn("starting http server on ", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("http server: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutting down http server")

		shutDownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutDownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown http server: %w", err)
		}

		h.log.Warn("http server shutdown completed")
	}

	return nil
}
