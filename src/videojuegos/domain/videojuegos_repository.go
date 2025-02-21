package domain

import "mi-api-go/src/videojuegos/domain/entities"

type IVideojuego interface {
	Save(videojuego entities.Videojuego) error
	Update(id int, videojuego entities.Videojuego) error
	Delete(id int) error
	FindByID(id int) (entities.Videojuego, error)
	GetAll() ([]entities.Videojuego, error)
	GetByStatus(deleted bool) ([]entities.Videojuego, error)
}
