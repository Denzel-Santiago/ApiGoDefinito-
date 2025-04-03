package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type CreateUsuarioUseCase struct {
	db              domain.IUsuario
	passwordService PasswordService
}

// Modificamos el constructor para recibir también el PasswordService
func NewCreateUsuarioUseCase(db domain.IUsuario, passwordService PasswordService) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{
		db:              db,
		passwordService: passwordService,
	}
}

func (uc *CreateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	// Si el usuario tiene una contraseña asignada, la codificamos
	if usuario.GetPassword() != "" {
		hashedPassword, err := uc.passwordService.Hash(usuario.GetPassword())
		if err != nil {
			return nil, err
		}
		usuario.SetPassword(hashedPassword)
	}

	err := uc.db.Save(*usuario)
	return usuario, err
}
