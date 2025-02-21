package application

import "mi-api-go/src/videojuegos/domain"

type DeleteVideojuegoUseCase struct {
	db domain.IVideojuego
}

func NewDeleteVideojuegoUseCase(db domain.IVideojuego) *DeleteVideojuegoUseCase {
	return &DeleteVideojuegoUseCase{
		db: db,
	}
}

func (uc *DeleteVideojuegoUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
