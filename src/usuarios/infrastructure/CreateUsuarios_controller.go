package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/usuarios/application"
	"mi-api-go/src/usuarios/domain/entities"
)

type CreateUsuarioController struct {
	CreateUsuarioUseCase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(createUsuarioUseCase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{
		CreateUsuarioUseCase: createUsuarioUseCase,
	}
}

func (ctrl *CreateUsuarioController) Run(c *gin.Context) {
	var usuario entities.Usuario

	if errJSON := c.ShouldBindJSON(&usuario); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del usuario inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	usuarioCreado, errAdd := ctrl.CreateUsuarioUseCase.Run(&usuario)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el usuario",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "El usuario ha sido agregado",
		"usuario": usuarioCreado,
	})
}
