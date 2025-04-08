package infrastructure

import (
	"mi-api-go/src/usuarios/application"
)

func InitUsuarioDependencies() (
	*CreateUsuarioController,
	*ViewUsuarioController,
	*UpdateUsuarioController,
	*DeleteUsuarioController,
	*ViewAllUsuariosController,
	*GetUsuariosByStatusController,
) {
	repo := NewMysqlUsuarioRepository()

	createUseCase := application.NewCreateUsuarioUseCase(repo)
	updateUseCase := application.NewUpdateUsuario(repo)
	viewUseCase := application.NewViewUsuario(repo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(repo)
	viewAllUseCase := application.NewViewAllUsuarios(repo)
	getUsuariosBystatus := application.NewGetUsuariosByStatusUseCase(repo)

	createController := NewCreateUsuarioController(createUseCase)
	viewController := NewViewUsuarioController(viewUseCase)
	updateController := NewUpdateUsuarioController(updateUseCase)
	deleteController := NewDeleteUsuarioController(deleteUseCase)
	viewAllController := NewViewAllUsuariosController(viewAllUseCase)
	getUsuariosBystatuscontroller := NewGetUsuariosByStatusController(getUsuariosBystatus)

	return createController, viewController, updateController, deleteController, viewAllController, getUsuariosBystatuscontroller
}
