package cifractx

import (
	"context"
	"net/http"
)

// MiddlewareWithContext add value to context with key
func MiddlewareWithContext(key ContextKey, value interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), key, value)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
