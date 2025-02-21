package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type UpdateUsuario struct {
	db domain.IUsuario
}

func NewUpdateUsuario(db domain.IUsuario) *UpdateUsuario {
	return &UpdateUsuario{db: db}
}

func (ue *UpdateUsuario) Execute(id int, usuario entities.Usuario) error {
	return ue.db.Update(id, usuario)
}
