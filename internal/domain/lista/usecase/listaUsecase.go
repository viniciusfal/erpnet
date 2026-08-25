package lista

import (
	"context"

	"github.com/google/uuid"
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

func (u *Usecase) Edit(ctx context.Context, input *lista.Lista) (*lista.Lista, error) {
	listaAtual, err := u.repo.FindById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	listaNova, err := listaAtual.Editar(input)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, listaNova); err != nil {
		return nil, err
	}

	return listaNova, nil
}

func (u *Usecase) Remove(ctx context.Context, idLista uuid.UUID) error {
	var l lista.Lista

	id, err := l.Remove(idLista)
	if err != nil {
		return err
	}

	_, err = u.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = u.repo.Delete(ctx, idLista)

	return err
}
