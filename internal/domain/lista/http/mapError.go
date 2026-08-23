package lista

import (
	"errors"
	"net/http"

	"github.com/viniciusfal/erpnet/internal/apierror"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

func mapError(err error) *apierror.Problem {
	switch {
	case errors.Is(err, lista.ERRNOMEOBRIGATORIO),
		errors.Is(err, lista.ERRIDOBRIGATORIO):
		return apierror.New(http.StatusBadRequest, "Erro de validacao", err.Error())

	default:
		return apierror.New(http.StatusInternalServerError, "Erro interno do servidor", "")
	}
}
