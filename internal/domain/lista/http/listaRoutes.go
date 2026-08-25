package lista

import "github.com/gin-gonic/gin"

func (c *ListaController) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/lista", c.Create)
	r.PUT("/lista/:id", c.Edit)
	r.GET("/lista", c.FindAll)
	r.DELETE("lista/:id", c.Remove)
}
