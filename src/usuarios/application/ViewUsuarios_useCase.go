package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type ViewUsuario struct {
	db domain.IUsuario
}

func NewViewUsuario(db domain.IUsuario) *ViewUsuario {
	return &ViewUsuario{db: db}
}

func (vu *ViewUsuario) Execute(id int) (entities.Usuario, error) {
	return vu.db.FindByID(id)
}
