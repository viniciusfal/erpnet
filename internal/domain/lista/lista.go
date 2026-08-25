package lista

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ERRIDOBRIGATORIO             = errors.New("O ID é obrigatorio")
	ERRNOMEOBRIGATORIO           = errors.New("O nome é obrigatorio")
	ERRRNOMECATEGORIAOBRIGATORIO = errors.New("O nome da categoria é obrigatoria")
)

type Lista struct {
	ID              uuid.UUID
	Nome            string
	Descricao       *string
	DataCriacao     time.Time
	DataAtualizacao *time.Time
}

type InputLista struct {
	Nome      string
	Descricao *string
}

func CreateLista(input InputLista) (*Lista, error) {
	idV7, err := uuid.NewV7()
	if err != nil {
		return nil, ERRIDOBRIGATORIO
	}

	if strings.TrimSpace(input.Nome) == "" {
		return nil, ERRNOMEOBRIGATORIO
	}

	l := Lista{
		ID:          idV7,
		Nome:        input.Nome,
		Descricao:   input.Descricao,
		DataCriacao: time.Now(),
	}

	return &l, nil
}

func (l *Lista) Editar(input *Lista) (*Lista, error) {

	if strings.TrimSpace(input.Nome) == "" {
		return nil, ERRNOMEOBRIGATORIO
	}

	atual := time.Now()

	novaLista := &Lista{
		ID:              l.ID,
		Nome:            input.Nome,
		Descricao:       input.Descricao,
		DataCriacao:     l.DataCriacao,
		DataAtualizacao: &atual,
	}

	return novaLista, nil
}
