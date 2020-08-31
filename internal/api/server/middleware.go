package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RequestIDKey struct{}

const requestIDKey = "request_id"

func (*RequestIDKey) String() string {
	return requestIDKey
}

func LogRequestMiddleware(log *zap.Logger) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				requestID := uuid.New().String()
				logger := log.With(
					zap.String("remote_addr", r.RemoteAddr),
					zap.String(requestIDKey, requestID),
				).Sugar()

				logger.Info(
					fmt.Sprintf("started %s %s", r.Method, r.RequestURI),
				)

				r = r.WithContext(
					context.WithValue(
						r.Context(), RequestIDKey{}, requestID,
					),
				)

				rw := &responseWriter{w, http.StatusOK}
				next.ServeHTTP(rw, r)

				msg := fmt.Sprintf(
					"completed with %d %s in %v",
					rw.code,
					http.StatusText(rw.code),
					time.Since(start),
				)

				switch {
				case rw.code >= 500:
					logger.Error(msg)
				case rw.code >= 400:
					logger.Warn(msg)
				default:
					logger.Info(msg)
				}
			},
		)
	}
}
