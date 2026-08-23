package lista

import "github.com/gin-gonic/gin"

func (c *ListaController) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/lista", c.Create)
	r.GET("/lista", c.FindAll)
}
