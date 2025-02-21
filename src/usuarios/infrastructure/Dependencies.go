// mi-api-go/src/usuarios/infraestructure/dependencies.go
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
	viewUseCase := application.NewViewUsuario(repo)
	updateUseCase := application.NewUpdateUsuario(repo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(repo)
	viewAllUseCase := application.NewViewAllUsuarios(repo)
	getUsuariosBystatus := application.NewGetUsuariosByStatusUseCase(repo)

	// Crear controladores
	createController := NewCreateUsuarioController(createUseCase)
	viewController := NewViewUsuarioController(viewUseCase)
	updateController := NewUpdateUsuarioController(updateUseCase)
	deleteController := NewDeleteUsuarioController(deleteUseCase)
	viewAllController := NewViewAllUsuariosController(viewAllUseCase)
	getUsuariosBystatuscontroller := NewGetUsuariosByStatusController(getUsuariosBystatus)

	return createController, viewController, updateController, deleteController, viewAllController, getUsuariosBystatuscontroller
}
