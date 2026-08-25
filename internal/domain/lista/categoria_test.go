package lista_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

func TestCategoria(t *testing.T) {

	payload := lista.InputCategoria{
		Nome: "categoria-01",
	}

	result, err := lista.CreateCategoria(payload)

	assert.NoError(t, err)
	assert.Equal(t, payload.Nome, result.Nome)
}
