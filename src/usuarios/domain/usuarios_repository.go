package domain

import "mi-api-go/src/usuarios/domain/entities"

type IUsuario interface {
	Save(usuario entities.Usuario) error
	Update(id int, usuario entities.Usuario) error
	Delete(id int) error
	FindByID(id int) (entities.Usuario, error)
	GetAll() ([]entities.Usuario, error)
	GetByStatus(deleted bool) ([]entities.Usuario, error)
}
