// mi-api-go/src/usuarios/infrastructure/dependencies.go
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
	passwordService := NewBcryptService() // Inicializamos el servicio de codificación

	createUseCase := application.NewCreateUsuarioUseCase(repo, passwordService) // Inyectamos el servicio de codificación
	updateUseCase := application.NewUpdateUsuario(repo, passwordService)        // Inyectamos el servicio de codificación
	viewUseCase := application.NewViewUsuario(repo)
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
