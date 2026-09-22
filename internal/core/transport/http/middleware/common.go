package core_http_middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/maximyunak/vetmessager/internal/core/auth"
	core_errors "github.com/maximyunak/vetmessager/internal/core/errors"
	core_logger "github.com/maximyunak/vetmessager/internal/core/logger"
	core_http_response "github.com/maximyunak/vetmessager/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-Id"

func RequestID() Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)
			next.ServeHTTP(w, r)
		})
	}
}

func CheckAuth(jwtManager *auth.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Authorization header is missing")
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Invalid authorization header")
				return
			}
			token, err := jwtManager.ParseToken(tokenString)
			if err != nil {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, "Unauthorized")
				return
			}

			ctx = context.WithValue(r.Context(), "uid", token.UserID)

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
				zap.String("method", r.Method),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			before := time.Now()
			rw := core_http_response.NewResponseWriter(w)

			log.Debug(
				">>> incoming http request",
				zap.Time("time", before.UTC()),
				zap.String("method", r.Method),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Now().Sub(before)),
			)
		})
	}
}
