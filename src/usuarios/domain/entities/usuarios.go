package entities

type Usuario struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	Apellido          string `json:"apellido"`
	Plataforma        string `json:"plataforma"`
	CorreoElectronico string `json:"correo_electronico"`
	Deleted           bool   `json:"deleted"`
}

// Constructor sin ID, ya que lo genera la BD
func NewUsuario(nombre, apellido, plataforma, correoElectronico string, deleted bool) *Usuario {
	return &Usuario{
		Nombre:            nombre,
		Apellido:          apellido,
		Plataforma:        plataforma,
		CorreoElectronico: correoElectronico,
		Deleted:           deleted,
	}
}

// Métodos Getters y Setters
func (u *Usuario) GetID() int {
	return u.ID
}

func (u *Usuario) SetID(id int) {
	u.ID = id
}

func (u *Usuario) GetNombre() string {
	return u.Nombre
}

func (u *Usuario) SetNombre(nombre string) {
	u.Nombre = nombre
}

func (u *Usuario) GetApellido() string {
	return u.Apellido
}

func (u *Usuario) SetApellido(apellido string) {
	u.Apellido = apellido
}

func (u *Usuario) GetPlataforma() string {
	return u.Plataforma
}

func (u *Usuario) SetPlataforma(plataforma string) {
	u.Plataforma = plataforma
}

func (u *Usuario) GetCorreoElectronico() string {
	return u.CorreoElectronico
}

func (u *Usuario) SetCorreoElectronico(correoElectronico string) {
	u.CorreoElectronico = correoElectronico
}

func (u *Usuario) IsDeleted() bool {
	return u.Deleted
}

func (u *Usuario) SetDeleted(deleted bool) {
	u.Deleted = deleted
}
