package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/usuarios/application"
)

type ViewAllUsuariosController struct {
	useCase *application.ViewAllUsuarios
}

func NewViewAllUsuariosController(useCase *application.ViewAllUsuarios) *ViewAllUsuariosController {
	return &ViewAllUsuariosController{useCase: useCase}
}

func (vuc *ViewAllUsuariosController) Execute(c *gin.Context) {
	usuarios, err := vuc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}
