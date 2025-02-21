package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/videojuegos/application"
	"mi-api-go/src/videojuegos/domain/entities"
)

type UpdateVideojuegoController struct {
	useCase *application.UpdateVideojuego
}

func NewUpdateVideojuegoController(useCase *application.UpdateVideojuego) *UpdateVideojuegoController {
	return &UpdateVideojuegoController{useCase: useCase}
}

func (uvc *UpdateVideojuegoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var videojuego entities.Videojuego
	if err := c.ShouldBindJSON(&videojuego); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uvc.useCase.Execute(id, videojuego)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el videojuego"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Videojuego actualizado exitosamente"})
}
