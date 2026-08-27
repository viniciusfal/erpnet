package lista_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
	"github.com/viniciusfal/erpnet/internal/domain/lista/mocks"
	usecase "github.com/viniciusfal/erpnet/internal/domain/lista/usecase"
)

func TestCategoria_Save(t *testing.T) {
	t.Run("cria e salva uma categoria com sucesso", func(t *testing.T) {
		repoMock := mocks.NewCategoriaRepository(t)
		repoMock.EXPECT().
			Save(mock.Anything, mock.AnythingOfType("*lista.Categoria")).
			Return(nil)

		uc := usecase.NewCategoriaUseCase(repoMock)

		c, err := uc.Create(context.Background(), lista.InputCategoria{
			Nome: "categoria-01",
		})

		assert.NoError(t, err)
		assert.NotNil(t, c)
	})

	t.Run("nao salva se domain rejeitar", func(t *testing.T) {
		repoMock := mocks.NewCategoriaRepository(t)

		svc := usecase.NewCategoriaUseCase(repoMock)

		_, err := svc.Create(context.Background(), lista.InputCategoria{
			// nome omitido de proposito
		})

		assert.ErrorIs(t, err, lista.ERRRNOMECATEGORIAOBRIGATORIO)
	})
}
