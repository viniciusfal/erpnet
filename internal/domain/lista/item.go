package lista

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/viniciusfal/erpnet/pkg"
)

var (
	ERRRNOMEITEMOBRIGATORIO               = errors.New("O nome do item é obrigatorio")
	ERRVALORDOITEMEOBRIGATORIO            = errors.New("O valor do item é obrigatorio")
	ERRSELECIONARUMACATEGORIAEOBRIGATORIO = errors.New("Selecione acategoria do seu item")
	ERRURLDOITEMESTACOMFORMATOINVALIDO    = errors.New("A URL do item esta com um formato invalido")
)

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
