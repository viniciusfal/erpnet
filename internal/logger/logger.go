// Loggers (O que aconteceu?)
// internal/logger/logger.go
// internal/logger/logger.go
package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func New() *slog.Logger {
	// dev: colorido e legível
	return slog.New(tint.NewTextHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))
}
