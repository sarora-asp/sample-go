package middleware

import (
	"context"
	"net/http"
)

func AuthContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "http_header", r.Header.Get("Authorization"))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
