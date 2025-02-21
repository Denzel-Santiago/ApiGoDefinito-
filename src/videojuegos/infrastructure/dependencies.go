package infrastructure

import (
	"mi-api-go/src/videojuegos/application"
)

func InitVideojuegoDependencies() (
	*CreateVideojuegoController,
	*ViewVideojuegoController,
	*UpdateVideojuegoController,
	*DeleteVideojuegoController,
	*ViewAllVideojuegosController,
	*GetVideojuegosByStatusController,
) {

	repo := NewMysqlVideojuegoRepository()

	createUseCase := application.NewCreateVideojuegoUseCase(repo)
	viewUseCase := application.NewViewVideojuego(repo)
	updateUseCase := application.NewUpdateVideojuego(repo)
	deleteUseCase := application.NewDeleteVideojuegoUseCase(repo)
	viewAllUseCase := application.NewViewAllVideojuegos(repo)
	getVideojuegosByStatus := application.NewGetVideojuegosByStatusUseCase(repo)

	// Crear controladores
	createController := NewCreateVideojuegoController(createUseCase)
	viewController := NewViewVideojuegoController(viewUseCase)
	updateController := NewUpdateVideojuegoController(updateUseCase)
	deleteController := NewDeleteVideojuegoController(deleteUseCase)
	viewAllController := NewViewAllVideojuegosController(viewAllUseCase)
	getVideojuegosByStatusController := NewGetVideojuegosByStatusController(getVideojuegosByStatus)

	return createController, viewController, updateController, deleteController, viewAllController, getVideojuegosByStatusController
}
