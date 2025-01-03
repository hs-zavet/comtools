package cifractx

import (
	"context"
	"fmt"
)

type ContextKey string

// WithValue adds a key-value pair to the context.
func WithValue(ctx context.Context, key ContextKey, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetValue returns the value for the key from the context.
func GetValue[T any](ctx context.Context, key ContextKey) (T, error) {
	val, ok := ctx.Value(key).(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("failed to get value from context for key: %v", key)
	}
	return val, nil
}
