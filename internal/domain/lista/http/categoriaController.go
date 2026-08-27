package lista

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erpnet/internal/apierror"
	"github.com/viniciusfal/erpnet/internal/domain/lista"
	usecase "github.com/viniciusfal/erpnet/internal/domain/lista/usecase"
	"github.com/viniciusfal/erpnet/internal/middleware"
)

type CategoriaController struct {
	uc *usecase.CategoriaUsecase
}

func NewCategoriaController(uc *usecase.CategoriaUsecase) *CategoriaController {
	return &CategoriaController{
		uc: uc,
	}
}

func (c *CategoriaController) Create(ctx *gin.Context) {
	log := middleware.LoggerFromContext(ctx.Request.Context())

	var req CategoriaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("Falha na Criacao da categoria", "error", err)
		payload := apierror.New(http.StatusBadRequest, "Payload invalido", err.Error())
		apierror.Respond(ctx, log, payload)
		return
	}

	category, err := c.uc.Create(ctx.Request.Context(), lista.InputCategoria(req))
	if err != nil {

		log.Error("Falha na Criacao da lista", "error", err)
		apierror.Respond(ctx, log, mapError(err))
		return
	}

	ctx.JSON(http.StatusCreated, ToResponseCategory(category))
}
