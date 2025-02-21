package application

import (
	"mi-api-go/src/videojuegos/domain"
	"mi-api-go/src/videojuegos/domain/entities"
)

type ViewVideojuego struct {
	db domain.IVideojuego
}

func NewViewVideojuego(db domain.IVideojuego) *ViewVideojuego {
	return &ViewVideojuego{db: db}
}

func (vv *ViewVideojuego) Execute(id int) (entities.Videojuego, error) {
	return vv.db.FindByID(id)
}
