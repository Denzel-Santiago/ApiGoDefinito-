package application

import (
	"mi-api-go/src/videojuegos/domain"
	"mi-api-go/src/videojuegos/domain/entities"
)

type UpdateVideojuego struct {
	db domain.IVideojuego
}

func NewUpdateVideojuego(db domain.IVideojuego) *UpdateVideojuego {
	return &UpdateVideojuego{db: db}
}

func (ue *UpdateVideojuego) Execute(id int, videojuego entities.Videojuego) error {
	return ue.db.Update(id, videojuego)
}
