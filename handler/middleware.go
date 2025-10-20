package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/google/uuid"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/otel"
)

func (handler *Handler) StandardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "X-Request-ID", uuid.NewString())

		lang := r.Header.Get("Lang")
		if lang == "" {
			lang = "ID"
		}

		ctx = context.WithValue(ctx, "RequestLang", lang)

		m := httpsnoop.CaptureMetrics(handler.PanicMiddlewares(next), w, r.WithContext(ctx))

		logger.Info(ctx, "http api request", map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL,
			"status":   m.Code,
			"duration": fmt.Sprint(m.Duration.Milliseconds(), "ms"),
		})
	})
}

func (handler *Handler) OtelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer.Start(r.Context(), "roll")
		defer func() {
			span.End()
		}()

		newRequest := r.Clone(ctx)

		next.ServeHTTP(w, newRequest)

	})
}

func (handler *Handler) PanicMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error(r.Context(), "panic occured", map[string]interface{}{
					"error": rec,
				})
				handler.WriteError(r.Context(), w, lib.ErrorInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (handler *Handler) AuthorizeSessionTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handler.WriteError(r.Context(), w, lib.ErrorInvalidToken)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		session, err := handler.BackendSkeleton.Usecase.ValidateToken(r.Context(), tokenString)
		if err != nil {
			handler.WriteError(r.Context(), w, lib.ErrorForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "CurrentUserID", session.UserID)
		ctx = context.WithValue(r.Context(), "CurrentSession", session)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
