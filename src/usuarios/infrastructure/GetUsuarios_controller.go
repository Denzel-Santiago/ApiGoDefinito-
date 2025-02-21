package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "mi-api-go/src/usuarios/application"
)

type GetUsuariosByStatusController struct {
	getByStatusUseCase *application.GetUsuariosByStatusUseCase
}

func NewGetUsuariosByStatusController(getByStatusUseCase *application.GetUsuariosByStatusUseCase) *GetUsuariosByStatusController {
	return &GetUsuariosByStatusController{
		getByStatusUseCase: getByStatusUseCase,
	}
}

func (ctrl *GetUsuariosByStatusController) Run(c *gin.Context) {
	statusParam := c.Param("status")
	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Estado inválido",
			"error":   err.Error(),
		})
		return
	}

	usuarios, err := ctrl.getByStatusUseCase.Run(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener usuarios por estado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"usuarios": usuarios,
	})
}
