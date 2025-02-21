package application

import (
	"mi-api-go/src/videojuegos/domain"
	"mi-api-go/src/videojuegos/domain/entities"
)

type CreateVideojuegoUseCase struct {
	db domain.IVideojuego
}

func NewCreateVideojuegoUseCase(db domain.IVideojuego) *CreateVideojuegoUseCase {
	return &CreateVideojuegoUseCase{
		db: db,
	}
}

func (uc *CreateVideojuegoUseCase) Run(videojuego *entities.Videojuego) (*entities.Videojuego, error) {
	err := uc.db.Save(*videojuego)
	return videojuego, err
}
