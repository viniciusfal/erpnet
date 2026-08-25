package lista_test

import (
	"testing"

	"github.com/google/uuid"
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
