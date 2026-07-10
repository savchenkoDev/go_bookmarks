package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "requestID"
const RequestIDHeader = "X-Request-ID"
type requestIDContextKey struct{}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set(RequestIDKey, requestID)
		c.Header(RequestIDHeader, requestID)
		ctx := context.WithValue(c.Request.Context(), requestIDContextKey{}, requestID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func RequestIDFromContext(c *gin.Context) string {
	if id, ok := c.Get(RequestIDKey); ok {
		if s, ok := id.(string); ok {
			return s
		}
	}
	return ""
}

func RequestIDFromGoContext(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDContextKey{}).(string); ok {
		return id
	}
	return ""
}