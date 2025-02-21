package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/usuarios/application"
	"mi-api-go/src/usuarios/domain/entities"
)

type UpdateUsuarioController struct {
	useCase *application.UpdateUsuario
}

func NewUpdateUsuarioController(useCase *application.UpdateUsuario) *UpdateUsuarioController {
	return &UpdateUsuarioController{useCase: useCase}
}

func (uuc *UpdateUsuarioController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var usuario entities.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uuc.useCase.Execute(id, usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}
