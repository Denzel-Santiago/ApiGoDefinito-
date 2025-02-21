package entities

type Videojuego struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Genero      string  `json:"genero"`
	Plataforma  string  `json:"plataforma"`
	Precio      float64 `json:"precio"`
	Deleted     bool    `json:"deleted"`
}

func NewVideojuego(id int, nombre, descripcion, genero, plataforma string, precio float64, deleted bool) *Videojuego {
	return &Videojuego{
		ID:          id,
		Nombre:      nombre,
		Descripcion: descripcion,
		Genero:      genero,
		Plataforma:  plataforma,
		Precio:      precio,
		Deleted:     deleted,
	}
}

func (v *Videojuego) GetID() int {
	return v.ID
}

func (v *Videojuego) SetID(id int) {
	v.ID = id
}

func (v *Videojuego) GetNombre() string {
	return v.Nombre
}

func (v *Videojuego) SetNombre(nombre string) {
	v.Nombre = nombre
}

func (v *Videojuego) GetDescripcion() string {
	return v.Descripcion
}

func (v *Videojuego) SetDescripcion(descripcion string) {
	v.Descripcion = descripcion
}

func (v *Videojuego) GetGenero() string {
	return v.Genero
}

func (v *Videojuego) SetGenero(genero string) {
	v.Genero = genero
}

func (v *Videojuego) GetPlataforma() string {
	return v.Plataforma
}

func (v *Videojuego) SetPlataforma(plataforma string) {
	v.Plataforma = plataforma
}

func (v *Videojuego) GetPrecio() float64 {
	return v.Precio
}

func (v *Videojuego) SetPrecio(precio float64) {
	v.Precio = precio
}

func (v *Videojuego) IsDeleted() bool {
	return v.Deleted
}

func (v *Videojuego) SetDeleted(deleted bool) {
	v.Deleted = deleted
}
