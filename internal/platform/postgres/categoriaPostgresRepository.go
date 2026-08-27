package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

type CategoriaRepository struct {
	db *pgxpool.Pool
}

func NewCategoriaRepository(db *pgxpool.Pool) *CategoriaRepository {
	return &CategoriaRepository{
		db: db,
	}
}

func (r *CategoriaRepository) Save(ctx context.Context, c *lista.Categoria) error {

	query := `
        INSERT INTO categoria (nome)
        VALUES ($1)
		RETURNING id
    `

	err := r.db.QueryRow(ctx, query, c.Nome).Scan(&c.ID)

	return err
}

func (r *CategoriaRepository) Update(ctx context.Context, categoria *lista.Categoria) error {
	return nil
}

func (r *CategoriaRepository) FindById(ctx context.Context, idCategoria int64) (*lista.Categoria, error) {
	return nil, nil
}

func (r *CategoriaRepository) FindAll(ctx context.Context) ([]lista.Categoria, error) {
	return nil, nil
}

func (r *CategoriaRepository) Delete(ctx context.Context, idCategoria int64) error {
	return nil
}
