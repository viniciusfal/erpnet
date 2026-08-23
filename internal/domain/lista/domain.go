package lista

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/viniciusfal/erpnet/pkg"
)

var (
	ERRIDOBRIGATORIO                      = errors.New("O ID é obrigatorio")
	ERRNOMEOBRIGATORIO                    = errors.New("O nome é obrigatorio")
	ERRRNOMECATEGORIAOBRIGATORIO          = errors.New("O nome da categoria é obrigatoria")
	ERRRNOMEITEMOBRIGATORIO               = errors.New("O nome do item é obrigatorio")
	ERRVALORDOITEMEOBRIGATORIO            = errors.New("O valor do item é obrigatorio")
	ERRSELECIONARUMACATEGORIAEOBRIGATORIO = errors.New("Selecione acategoria do seu item")
	ERRURLDOITEMESTACOMFORMATOINVALIDO    = errors.New("A URL do item esta com um formato invalido")
)

type Lista struct {
	ID              uuid.UUID
	Nome            string
	Descricao       *string
	DataCriacao     time.Time
	DataAtualizacao *time.Time
}

type Item struct {
	ID              uuid.UUID
	Nome            string
	Descricao       *string
	Quantidade      int
	UrlImage        *string
	CategoriaID     int64
	ValorUnitario   decimal.Decimal
	DataCriacao     time.Time
	DataAtualizacao *time.Time
}

type Categoria struct {
	ID   int64
	Nome string
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

	if input.Nome == "" {
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

type InputCategoria struct {
	Nome string
}

func CreateCategoria(input InputCategoria) (*Categoria, error) {
	if input.Nome == "" {
		return nil, ERRRNOMECATEGORIAOBRIGATORIO
	}

	c := Categoria{
		Nome: input.Nome,
	}

	return &c, nil
}

type InputItem struct {
	Nome          string
	Descricao     *string
	Quantidade    int
	UrlImage      *string
	ValorUnitario decimal.Decimal
	CategoriaId   int64
}

func CreateItem(input InputItem) (*Item, error) {
	idV7, err := uuid.NewV7()
	if err != nil {
		return nil, ERRIDOBRIGATORIO
	}

	if input.Quantidade < 1 {
		input.Quantidade = 1
	}

	switch {
	case input.Nome == "":
		return nil, ERRRNOMEITEMOBRIGATORIO

	case input.ValorUnitario.IsZero():
		return nil, ERRVALORDOITEMEOBRIGATORIO

	case input.CategoriaId == 0:
		return nil, ERRSELECIONARUMACATEGORIAEOBRIGATORIO

	case input.UrlImage != nil && !pkg.IsValidURL(*input.UrlImage):
		return nil, ERRURLDOITEMESTACOMFORMATOINVALIDO
	}

	i := Item{
		ID:            idV7,
		Nome:          input.Nome,
		Descricao:     input.Descricao,
		Quantidade:    input.Quantidade,
		UrlImage:      input.UrlImage,
		ValorUnitario: input.ValorUnitario,
		CategoriaID:   input.CategoriaId,
		DataCriacao:   time.Now(),
	}

	return &i, nil
}
