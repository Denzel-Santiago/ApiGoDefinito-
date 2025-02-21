package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "mi-api-go/src/videojuegos/application"
)

type DeleteVideojuegoController struct {
	deleteUseCase *application.DeleteVideojuegoUseCase
}

func NewDeleteVideojuegoController(deleteUseCase *application.DeleteVideojuegoUseCase) *DeleteVideojuegoController {
	return &DeleteVideojuegoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteVideojuegoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el videojuego",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Videojuego eliminado exitosamente",
	})
}
