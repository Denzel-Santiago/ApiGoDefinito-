package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"mi-api-go/src/core"
	"mi-api-go/src/usuarios/domain"
	"mi-api-go/src/usuarios/domain/entities"
)

type MysqlUsuario struct {
	conn *sql.DB
}

func NewMysqlUsuarioRepository() domain.IUsuario {
	conn := core.GetDB()
	return &MysqlUsuario{conn: conn}
}

func (mysql *MysqlUsuario) Save(usuario entities.Usuario) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO usuarios (nombre, apellido, plataforma, correo_electronico, deleted) VALUES (?, ?, ?, ?, ?)",
		usuario.Nombre,
		usuario.Apellido,
		usuario.Plataforma,
		usuario.CorreoElectronico,
		usuario.Deleted,
	)
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	usuario.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlUsuario) Update(id int, usuario entities.Usuario) error {
	result, err := mysql.conn.Exec(
		"UPDATE usuarios SET nombre = ?, apellido = ?, plataforma = ?, correo_electronico = ?, deleted = ? WHERE id = ?",
		usuario.Nombre,
		usuario.Apellido,
		usuario.Plataforma,
		usuario.CorreoElectronico,
		usuario.Deleted,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el usuario con ID:", id)
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlUsuario) Delete(id int) error {
	_, err := mysql.conn.Exec("UPDATE usuarios SET deleted = true WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar (soft delete) el usuario:", err)
		return err
	}
	return nil
}

func (mysql *MysqlUsuario) FindByID(id int) (entities.Usuario, error) {
	var usuario entities.Usuario
	row := mysql.conn.QueryRow("SELECT id, nombre, apellido, plataforma, correo_electronico, deleted FROM usuarios WHERE id = ?", id)

	err := row.Scan(
		&usuario.ID,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.Plataforma,
		&usuario.CorreoElectronico,
		&usuario.Deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuario no encontrado:", err)
			return entities.Usuario{}, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el usuario por ID:", err)
		return entities.Usuario{}, err
	}

	return usuario, nil
}

func (mysql *MysqlUsuario) GetAll() ([]entities.Usuario, error) {
	var usuarios []entities.Usuario

	rows, err := mysql.conn.Query("SELECT id, nombre, apellido, plataforma, correo_electronico, deleted FROM usuarios")
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usuario entities.Usuario
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nombre,
			&usuario.Apellido,
			&usuario.Plataforma,
			&usuario.CorreoElectronico,
			&usuario.Deleted,
		)
		if err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return usuarios, nil
}

func (mysql *MysqlUsuario) GetByStatus(deleted bool) ([]entities.Usuario, error) {
	var usuarios []entities.Usuario

	rows, err := mysql.conn.Query("SELECT id, nombre, apellido, plataforma, correo_electronico, deleted FROM usuarios WHERE deleted = ?", deleted)
	if err != nil {
		log.Println("Error al obtener usuarios por estado:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usuario entities.Usuario
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nombre,
			&usuario.Apellido,
			&usuario.Plataforma,
			&usuario.CorreoElectronico,
			&usuario.Deleted,
		)
		if err != nil {
			log.Println("Error al filtrar a los usuarios:", err)
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al filtrar los usuarios:", err)
		return nil, err
	}

	return usuarios, nil
}
