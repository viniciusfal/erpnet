package lista

import (
	"time"

	"github.com/google/uuid"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
)

type CreteListaReq struct {
	Nome      string  `json:"nome" binding:"required,min=3,max=100"`
	Descricao *string `json:"descricao" binding:"omitempty"`
}

type ListaResp struct {
	ID              uuid.UUID  `json:"id"`
	Nome            string     `json:"nome"`
	Descricao       *string    `json:"descricao"`
	DataCriacao     time.Time  `json:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao"`
}

func ToResponse(l *lista.Lista) ListaResp {
	return ListaResp{
		ID:              l.ID,
		Nome:            l.Nome,
		Descricao:       l.Descricao,
		DataCriacao:     l.DataCriacao,
		DataAtualizacao: l.DataAtualizacao,
	}
}

func ToResponseList(listas []lista.Lista) []ListaResp {
	resp := make([]ListaResp, 0, len(listas)) // Garante slice vazio [] em vez de nil no JSON
	for _, l := range listas {
		resp = append(resp, ToResponse(&l)) // Ou a conversão individual
	}
	return resp
}
