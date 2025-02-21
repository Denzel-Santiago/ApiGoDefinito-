package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/videojuegos/application"
)

type ViewVideojuegoController struct {
	useCase *application.ViewVideojuego
}

func NewViewVideojuegoController(useCase *application.ViewVideojuego) *ViewVideojuegoController {
	return &ViewVideojuegoController{useCase: useCase}
}

func (vvc *ViewVideojuegoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	videojuego, err := vvc.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Videojuego no encontrado"})
		return
	}

	c.JSON(http.StatusOK, videojuego)
}
