package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type UpdateUsuario struct {
	db              domain.IUsuario
	passwordService PasswordService
}

func NewUpdateUsuario(db domain.IUsuario, passwordService PasswordService) *UpdateUsuario {
	return &UpdateUsuario{
		db:              db,
		passwordService: passwordService,
	}
}

func (ue *UpdateUsuario) Execute(id int, usuario entities.Usuario) error {
	// Si el usuario tiene una contraseña asignada, la codificamos
	if usuario.GetPassword() != "" {
		hashedPassword, err := ue.passwordService.Hash(usuario.GetPassword())
		if err != nil {
			return err
		}
		usuario.SetPassword(hashedPassword)
	}

	return ue.db.Update(id, usuario)
}
