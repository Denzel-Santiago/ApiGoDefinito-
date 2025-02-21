package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type CreateUsuarioUseCase struct {
	db domain.IUsuario
}

func NewCreateUsuarioUseCase(db domain.IUsuario) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{
		db: db,
	}
}

func (uc *CreateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	err := uc.db.Save(*usuario)
	return usuario, err
}
