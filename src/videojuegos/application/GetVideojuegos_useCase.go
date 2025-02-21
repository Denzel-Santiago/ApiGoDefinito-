package application

import (
	"mi-api-go/src/videojuegos/domain"
	"mi-api-go/src/videojuegos/domain/entities"
)

type GetVideojuegosByStatusUseCase struct {
	db domain.IVideojuego
}

func NewGetVideojuegosByStatusUseCase(db domain.IVideojuego) *GetVideojuegosByStatusUseCase {
	return &GetVideojuegosByStatusUseCase{
		db: db,
	}
}

func (uc *GetVideojuegosByStatusUseCase) Run(status bool) ([]entities.Videojuego, error) {
	return uc.db.GetByStatus(status)
}
