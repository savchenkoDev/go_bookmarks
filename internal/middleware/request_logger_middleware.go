// internal/middleware/request_logger_middleware.go
package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return RequestLoggerWith(slog.Default())
}

func RequestLoggerWith(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := c.Writer.Status()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		attrs := []any{
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("query", c.Request.URL.RawQuery),
			slog.Int("status", status),
			slog.Int("bytes", c.Writer.Size()),
			slog.Duration("latency", time.Since(start)),
			slog.String("client_ip", c.ClientIP()),
		}

		if userID, ok := c.Get("userID"); ok {
			attrs = append(attrs, slog.Int64("user_id", userID.(int64)))
		}

		switch {
		case status >= 500:
			log.ErrorContext(c.Request.Context(), "http_request", attrs...)
		case status >= 400:
			log.WarnContext(c.Request.Context(), "http_request", attrs...)
		default:
			log.InfoContext(c.Request.Context(), "http_request", attrs...)
		}
	}
}
