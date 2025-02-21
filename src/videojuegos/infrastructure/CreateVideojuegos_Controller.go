package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/videojuegos/application"
	"mi-api-go/src/videojuegos/domain/entities"
)

type CreateVideojuegoController struct {
	CreateVideojuegoUseCase *application.CreateVideojuegoUseCase
}

func NewCreateVideojuegoController(createVideojuegoUseCase *application.CreateVideojuegoUseCase) *CreateVideojuegoController {
	return &CreateVideojuegoController{
		CreateVideojuegoUseCase: createVideojuegoUseCase,
	}
}

func (ctrl *CreateVideojuegoController) Run(c *gin.Context) {
	var videojuego entities.Videojuego

	if errJSON := c.ShouldBindJSON(&videojuego); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del videojuego inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	videojuegoCreado, errAdd := ctrl.CreateVideojuegoUseCase.Run(&videojuego)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el videojuego",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "El videojuego ha sido agregado",
		"videojuego": videojuegoCreado,
	})
}
