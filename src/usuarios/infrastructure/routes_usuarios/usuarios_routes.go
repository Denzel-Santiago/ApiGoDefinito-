package routesusuarios

import (
	"github.com/gin-gonic/gin"
	"mi-api-go/src/usuarios/infrastructure" // Se importa correctamente la dependencia
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (router *Router) Run() {
	// Se llama a InitUsuarioDependencies con el prefijo correcto
	createController, viewController, updateController, deleteController, viewAllController, getUsuariosByStatusController := infrastructure.InitUsuarioDependencies()

	usuarioGroup := router.engine.Group("/usuarios")
	{
		usuarioGroup.POST("/", createController.Run)
		usuarioGroup.GET("/:id", viewController.Execute)
		usuarioGroup.PUT("/:id", updateController.Execute)
		usuarioGroup.DELETE("/:id", deleteController.Run)
		usuarioGroup.GET("/", viewAllController.Execute)
		usuarioGroup.GET("/status/:status", getUsuariosByStatusController.Run)
	}
}
