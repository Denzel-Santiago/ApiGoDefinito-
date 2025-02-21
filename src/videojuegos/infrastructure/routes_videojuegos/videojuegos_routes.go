package routesvideojuegos

import (
	"github.com/gin-gonic/gin"
	"mi-api-go/src/videojuegos/infrastructure" // Ahora se usa correctamente
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
	// Llamada correcta a InitVideojuegoDependencies
	createController, viewController, updateController, deleteController, viewAllController, getVideojuegosByStatusController := infrastructure.InitVideojuegoDependencies()

	videojuegoGroup := router.engine.Group("/videojuegos")
	{
		videojuegoGroup.POST("/", createController.Run)
		videojuegoGroup.GET("/:id", viewController.Execute)
		videojuegoGroup.PUT("/:id", updateController.Execute)
		videojuegoGroup.DELETE("/:id", deleteController.Run)
		videojuegoGroup.GET("/", viewAllController.Execute)
		videojuegoGroup.GET("/status/:status", getVideojuegosByStatusController.Run)
	}
}
