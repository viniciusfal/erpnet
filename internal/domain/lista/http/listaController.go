package lista

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/viniciusfal/erpnet/internal/apierror"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
	usecase "github.com/viniciusfal/erpnet/internal/domain/lista/usecase"
	"github.com/viniciusfal/erpnet/internal/middleware"
)

type ListaController struct {
	uc *usecase.Usecase
}

func NewController(uc *usecase.Usecase) *ListaController {
	return &ListaController{
		uc: uc,
	}
}

func (c *ListaController) Create(ctx *gin.Context) {
	log := middleware.LoggerFromContext(ctx.Request.Context())

	var req ListaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("Falha na Criacao da lista", "error", err)
		payload := apierror.New(http.StatusBadRequest, "Payload invalido", err.Error())
		apierror.Respond(ctx, log, payload)
		return
	}

	lista, err := c.uc.Create(ctx.Request.Context(), lista.InputLista(req))
	if err != nil {

		log.Error("Falha na Criacao da lista", "error", err)
		apierror.Respond(ctx, log, mapError(err))
		return
	}

	ctx.JSON(http.StatusCreated, ToResponse(lista))
}

func (c *ListaController) Edit(ctx *gin.Context) {
	log := middleware.LoggerFromContext(ctx.Request.Context())

	id := ctx.Param("id")
	if strings.TrimSpace(id) == "" {
		payload := apierror.New(http.StatusBadRequest, "ID da lista é obrigatório", "")
		apierror.Respond(ctx, log, payload)
		return
	}

	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.Error("ID inválido", "error", err)
		payload := apierror.New(http.StatusBadRequest, "ID da lista é inválido", err.Error())
		apierror.Respond(ctx, log, payload)
		return
	}

	var req ListaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("Falha na edicao da lista", "error", err)
		payload := apierror.New(http.StatusBadRequest, "Payload invalido", err.Error())
		apierror.Respond(ctx, log, payload)
		return
	}

	input := lista.Lista{
		ID:        idParsed,
		Nome:      req.Nome,
		Descricao: req.Descricao,
	}

	listaAtualizada, err := c.uc.Edit(ctx.Request.Context(), &input)
	if err != nil {
		log.Error("Erro ao processar edicao da lista", "error", err)
		payload := apierror.New(http.StatusInternalServerError, "Erro ao editar lista", err.Error())
		apierror.Respond(ctx, log, payload)
		return
	}

	ctx.JSON(http.StatusOK, ToResponse(listaAtualizada))
}

func (c *ListaController) FindAll(ctx *gin.Context) {
	log := middleware.LoggerFromContext(ctx.Request.Context())

	listas, err := c.uc.FindAll(ctx.Request.Context())
	if err != nil {
		log.Error("Falha na Busca das listas", "error", err)
		apierror.Respond(ctx, log, mapError(err))
		return
	}

	ctx.JSON(http.StatusOK, ToResponseList(listas))
}

func (c *ListaController) Remove(ctx *gin.Context) {
	log := middleware.LoggerFromContext(ctx.Request.Context())

	id := ctx.Param("id")
	if strings.TrimSpace(id) == "" {
		payload := apierror.New(http.StatusBadRequest, "ID da lista é obrigatório", "")
		apierror.Respond(ctx, log, payload)
		return
	}

	idParsed, err := uuid.Parse(id)
	if err != nil {
		log.Error("ID inválido", "error", err)
		payload := apierror.New(http.StatusBadRequest, "ID da lista é inválido", err.Error())
		apierror.Respond(ctx, log, payload)
		return
	}

	if err = c.uc.Remove(ctx.Request.Context(), idParsed); err != nil {
		log.Error("ID inválido", "erro ao tentar remover lista", err)
		apierror.Respond(ctx, log, mapError(err))
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{id: "Registro excluido com sucesso"})

}
