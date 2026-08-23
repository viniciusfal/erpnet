package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, uri string) (*pgxpool.Pool, error) {
	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := pgxpool.New(c, uri)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao postgres: %w", err)
	}

	if err := client.Ping(c); err != nil {
		client.Close()
		return nil, fmt.Errorf("postgres não respondeu ao Ping: %w", err)
	}

	return client, nil
}

func Disconnect(client *pgxpool.Pool, log *slog.Logger) {
	if client == nil {
		return
	}

	if log != nil {
		log.Info("Encerrando pool de conexões com o Postgres...")
	}

	client.Close()
}
