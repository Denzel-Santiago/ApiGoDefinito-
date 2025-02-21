package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "mi-api-go/src/videojuegos/application"
)

type GetVideojuegosByStatusController struct {
	getByStatusUseCase *application.GetVideojuegosByStatusUseCase
}

func NewGetVideojuegosByStatusController(getByStatusUseCase *application.GetVideojuegosByStatusUseCase) *GetVideojuegosByStatusController {
	return &GetVideojuegosByStatusController{
		getByStatusUseCase: getByStatusUseCase,
	}
}

func (ctrl *GetVideojuegosByStatusController) Run(c *gin.Context) {
	statusParam := c.Param("status")
	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Estado inválido",
			"error":   err.Error(),
		})
		return
	}

	videojuegos, err := ctrl.getByStatusUseCase.Run(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener videojuegos por estado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"videojuegos": videojuegos,
	})
}
