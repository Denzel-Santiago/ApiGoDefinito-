package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/usuarios/application"
)

type ViewUsuarioController struct {
	useCase *application.ViewUsuario
}

func NewViewUsuarioController(useCase *application.ViewUsuario) *ViewUsuarioController {
	return &ViewUsuarioController{useCase: useCase}
}

func (vuc *ViewUsuarioController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := vuc.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
