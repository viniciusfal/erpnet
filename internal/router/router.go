package router

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erpnet/internal/health"
	"github.com/viniciusfal/erpnet/internal/middleware"
)

type Dependencies struct {
	Logger  *slog.Logger
	Health  *health.Health
	Modules []Module
}

type Module interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

func New(deps *Dependencies) *gin.Engine {
	r := gin.New()

	r.Use(
		middleware.Recovery(deps.Logger),
		middleware.StructuredLogger(deps.Logger),
	)

	r.GET("/health", health.New().Liveness)

	api := r.Group("/api")
	for _, m := range deps.Modules {
		m.RegisterRoutes(api)
	}

	return r
}
