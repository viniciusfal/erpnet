package lista_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
	"github.com/viniciusfal/erpnet/internal/domain/lista/mocks"
	usecase "github.com/viniciusfal/erpnet/internal/domain/lista/usecase"
)

func TestService_Save(t *testing.T) {
	t.Run("cria e salva uma lista com sucesso", func(t *testing.T) {
		repoMock := mocks.NewListaRepository(t)
		repoMock.EXPECT().
			Save(mock.Anything, mock.AnythingOfType("*lista.Lista")).
			Return(nil)

		uc := usecase.NewUseCase(repoMock)

		l, err := uc.Create(context.Background(), lista.InputLista{
			Nome: "lista-01",
		})

		assert.NoError(t, err)
		assert.NotNil(t, l)
	})

	t.Run("nao salva se domain rejeitar", func(t *testing.T) {
		repoMock := mocks.NewListaRepository(t)

		svc := usecase.NewUseCase(repoMock)

		_, err := svc.Create(context.Background(), lista.InputLista{
			// nome omitido de proposito
		})

		assert.ErrorIs(t, err, lista.ERRNOMEOBRIGATORIO)
	})
}

func TestService_FindAll(t *testing.T) {
	t.Run("tras uma collection de lista", func(t *testing.T) {
		idv7, _ := uuid.NewV7()

		repoMock := mocks.NewListaRepository(t)

		expectedList := []lista.Lista{
			{ID: idv7, Nome: "Lista Teste"},
		}

		repoMock.EXPECT().
			FindAll(mock.Anything).
			Return(expectedList, nil)

		uc := usecase.NewUseCase(repoMock)

		l, err := uc.FindAll(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, l)
		assert.Equal(t, expectedList, l)
	})
}
