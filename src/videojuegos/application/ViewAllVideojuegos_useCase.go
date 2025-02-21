package application

import (
	"mi-api-go/src/videojuegos/domain"
	"mi-api-go/src/videojuegos/domain/entities"
)

type ViewAllVideojuegos struct {
	db domain.IVideojuego
}

func NewViewAllVideojuegos(db domain.IVideojuego) *ViewAllVideojuegos {
	return &ViewAllVideojuegos{db: db}
}

func (ve *ViewAllVideojuegos) Execute() ([]entities.Videojuego, error) {
	return ve.db.GetAll()
}
