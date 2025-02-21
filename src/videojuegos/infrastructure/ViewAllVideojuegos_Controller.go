package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/videojuegos/application"
)

type ViewAllVideojuegosController struct {
	useCase *application.ViewAllVideojuegos
}

func NewViewAllVideojuegosController(useCase *application.ViewAllVideojuegos) *ViewAllVideojuegosController {
	return &ViewAllVideojuegosController{useCase: useCase}
}

func (vvc *ViewAllVideojuegosController) Execute(c *gin.Context) {
	videojuegos, err := vvc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los videojuegos"})
		return
	}

	c.JSON(http.StatusOK, videojuegos)
}
