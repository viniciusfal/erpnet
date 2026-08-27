package lista

import (
	"context"

	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

type CategoriaUsecase struct {
	repo lista.CategoriaRepository
}

func NewCategoriaUseCase(repo lista.CategoriaRepository) *CategoriaUsecase {
	return &CategoriaUsecase{
		repo: repo,
	}
}

func (u *CategoriaUsecase) Create(ctx context.Context,
	input lista.InputCategoria) (*lista.Categoria, error) {

	l, err := lista.CreateCategoria(input)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Save(ctx, l); err != nil {
		return nil, err
	}

	return l, nil
}
