package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

type ListaRepository struct {
	db *pgxpool.Pool
}

func NewListaRepository(db *pgxpool.Pool) *ListaRepository {
	return &ListaRepository{
		db: db,
	}
}

func (r *ListaRepository) Save(ctx context.Context, l *lista.Lista) error {
	query := `
        INSERT INTO lista (id, nome, descricao)
        VALUES ($1, $2, $3)
    `

	_, err := r.db.Exec(ctx, query, l.ID, l.Nome, l.Descricao)

	return err
}

func (r *ListaRepository) Update(ctx context.Context, l *lista.Lista) error {

	query := `
		UPDATE  lista
		SET 
			nome = $2
			, descricao = $3
			, data_atualizacao = $4
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		l.ID,
		l.Nome,
		l.Descricao,
		l.DataAtualizacao)

	return err
}

func (r *ListaRepository) FindById(ctx context.Context, idLista uuid.UUID) (*lista.Lista, error) {
	var l lista.Lista

	query := `
		SELECT 
			id
			, nome
			, descricao 
			, data_criacao
			, data_atualizacao
		FROM lista
		WHERE id = $1
	`

	if err := r.db.QueryRow(ctx, query, idLista).Scan(
		&l.ID,
		&l.Nome,
		&l.Descricao,
		&l.DataCriacao,
		&l.DataAtualizacao,
	); err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *ListaRepository) FindAll(ctx context.Context) ([]lista.Lista, error) {
	query := `
		SELECT 
			id
			, nome
			, descricao 
			, data_criacao
			, data_atualizacao
		FROM lista
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var listas []lista.Lista
	for rows.Next() {
		var l lista.Lista

		if err := rows.Scan(
			&l.ID,
			&l.Nome,
			&l.Descricao,
			&l.DataCriacao,
			&l.DataAtualizacao,
		); err != nil {
			return nil, err
		}

		listas = append(listas, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listas, nil
}

func (r *ListaRepository) Delete(ctx context.Context, idLista uuid.UUID) error {
	return nil
}
