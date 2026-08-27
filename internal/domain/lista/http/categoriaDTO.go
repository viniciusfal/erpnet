package lista

import (
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

type CategoriaReq struct {
	Nome string `json:"nome" binding:"required,min=3"`
}

type CategoriaResp struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}

func ToResponseCategory(c *lista.Categoria) CategoriaResp {
	return CategoriaResp{
		ID:   c.ID,
		Nome: c.Nome,
	}
}
