package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erpnet/internal/apierror"
)

type ctxKey string

const loggerCtxKey ctxKey = "logger"

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(loggerCtxKey).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}

func StructuredLogger(base *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqLogger := base.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)

		ctx := context.WithValue(c.Request.Context(), loggerCtxKey, reqLogger)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		reqLogger.Info(
			"request finalizada",
			slog.Int("status", c.Writer.Status()),
			slog.Duration("latency", time.Since(start)),
		)
	}
}

func Recovery(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recuperado", slog.Any("panic", r))
				p := apierror.New(http.StatusInternalServerError, "Erro interno do servidor", "")
				apierror.Respond(c, log, p)
			}
		}()
		c.Next()
	}
}
