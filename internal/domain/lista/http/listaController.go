package lista

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	var req CreteListaReq
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
