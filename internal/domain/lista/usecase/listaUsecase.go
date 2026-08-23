package lista

import (
	"context"

	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

type Usecase struct {
	repo lista.ListaRepository
}

func NewUseCase(repo lista.ListaRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context,
	input lista.InputLista) (*lista.Lista, error) {

	l, err := lista.CreateLista(input)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Save(ctx, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (u *Usecase) FindAll(ctx context.Context) ([]lista.Lista, error) {

	l, err := u.repo.FindAll(ctx)
	if err != nil {
		return []lista.Lista{}, err
	}

	return l, err
}
