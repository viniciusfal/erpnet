package lista_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

func TestLista(t *testing.T) {

	payload := lista.InputLista{
		Nome: "lista-01",
	}

	result, err := lista.CreateLista(payload)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, result.ID)
	assert.Equal(t, payload.Nome, result.Nome)
	assert.NotNil(t, result.ID)
	assert.Nil(t, result.DataAtualizacao)
}

func TestCategoria(t *testing.T) {

	payload := lista.InputCategoria{
		Nome: "categoria-01",
	}

	result, err := lista.CreateCategoria(payload)

	assert.NoError(t, err)
	assert.Equal(t, payload.Nome, result.Nome)
}

func TestItem(t *testing.T) {
	valor, _ := decimal.NewFromString("199.99")
	urlInvalida := "htto://url_invalida"

	testes := []struct {
		name    string
		input   lista.InputItem
		wantErr error
	}{
		{
			name: "criar um item valido",
			input: lista.InputItem{
				Nome:          "item-01",
				ValorUnitario: valor,
				CategoriaId:   1,
			},
		},
		{
			name:    "sem nome do item",
			input:   lista.InputItem{ValorUnitario: valor, CategoriaId: 1},
			wantErr: lista.ERRRNOMEITEMOBRIGATORIO,
		},
		{
			name:    "sem valor unitario",
			input:   lista.InputItem{Nome: "item-01", CategoriaId: 1},
			wantErr: lista.ERRVALORDOITEMEOBRIGATORIO,
		},
		{
			name:    "sem categoria do item",
			input:   lista.InputItem{Nome: "item-01", ValorUnitario: valor},
			wantErr: lista.ERRSELECIONARUMACATEGORIAEOBRIGATORIO,
		},
		{
			name: "Url invalida",
			input: lista.InputItem{
				Nome:          "item-01",
				ValorUnitario: valor,
				CategoriaId:   1,
				UrlImage:      &urlInvalida},
			wantErr: lista.ERRURLDOITEMESTACOMFORMATOINVALIDO,
		},
	}

	for _, tt := range testes {
		t.Run(tt.name, func(t *testing.T) {
			item, err := lista.CreateItem(tt.input)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.NotEqual(t, uuid.Nil, item.ID)
			assert.Equal(t, tt.input.Nome, item.Nome)
			assert.Equal(t, tt.input.CategoriaId, item.CategoriaID)
			assert.Equal(t, tt.input.ValorUnitario, item.ValorUnitario)
			assert.Equal(t, tt.input.UrlImage, item.UrlImage)
			assert.Equal(t, tt.input.Descricao, item.Descricao)
			assert.WithinDuration(t, time.Now(), item.DataCriacao, time.Second)
			assert.Nil(t, item.DataAtualizacao)
		})
	}
}
