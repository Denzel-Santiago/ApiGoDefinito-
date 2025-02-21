package application

import (
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type ViewAllUsuarios struct {
	db domain.IUsuario
}

func NewViewAllUsuarios(db domain.IUsuario) *ViewAllUsuarios {
	return &ViewAllUsuarios{db: db}
}

func (ve *ViewAllUsuarios) Execute() ([]entities.Usuario, error) {
	return ve.db.GetAll()
}
