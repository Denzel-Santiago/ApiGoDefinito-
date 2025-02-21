package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type GetUsuariosByStatusUseCase struct {
	db domain.IUsuario
}

func NewGetUsuariosByStatusUseCase(db domain.IUsuario) *GetUsuariosByStatusUseCase {
	return &GetUsuariosByStatusUseCase{
		db: db,
	}
}

func (uc *GetUsuariosByStatusUseCase) Run(status bool) ([]entities.Usuario, error) {
	return uc.db.GetByStatus(status)
}
