package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/viniciusfal/erpnet/internal/config"
	controller "github.com/viniciusfal/erpnet/internal/domain/lista/http"
	usecase "github.com/viniciusfal/erpnet/internal/domain/lista/usecase"
	"github.com/viniciusfal/erpnet/internal/health"
	"github.com/viniciusfal/erpnet/internal/logger"
	"github.com/viniciusfal/erpnet/internal/platform/postgres"
	"github.com/viniciusfal/erpnet/internal/router"
	"github.com/viniciusfal/erpnet/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("aviso: .env não encontrado, usando variáveis de ambiente do sistema")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	l := logger.New()

	conn, err := postgres.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	h := health.New()
	h.Register("postgres", func(ctx context.Context) error {
		return conn.Ping(ctx)
	})

	listaRepo := postgres.NewListaRepository(conn)
	listaUC := usecase.NewUseCase(listaRepo)
	listaCtrl := controller.NewController(listaUC)

	deps := &router.Dependencies{
		Logger: l,
		Health: h,
		Modules: []router.Module{
			listaCtrl,
		},
	}
	r := router.New(deps)

	if err := server.Run(ctx, r, cfg.Port, l); err != nil {
		l.Error("erro ao encerrar servidor", "error", err)
	}
}
