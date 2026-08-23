package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Checker func(ctx context.Context) error

type Health struct {
	checks map[string]Checker
}

func New() *Health {
	return &Health{checks: make(map[string]Checker)}
}

// Register adiciona uma dependência a ser verificada no readiness.
func (h *Health) Register(name string, check Checker) {
	h.checks[name] = check
}

// Liveness só confirma que o processo responde — sem checar dependências.
func (h *Health) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
