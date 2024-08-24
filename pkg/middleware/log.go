package middleware

import (
	"log/slog"
	"net/http"
)

type router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func LogMiddleware(r router) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Debug("request:", "protocol", r.Proto, "method", r.Method, "URL", r.URL, "remote addr", r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	}
}
