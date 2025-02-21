package application

import "mi-api-go/src/usuarios/domain"

type DeleteUsuarioUseCase struct {
	db domain.IUsuario
}

func NewDeleteUsuarioUseCase(db domain.IUsuario) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{
		db: db,
	}
}

func (uc *DeleteUsuarioUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
