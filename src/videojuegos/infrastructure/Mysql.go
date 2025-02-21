package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"mi-api-go/src/core"
	"mi-api-go/src/videojuegos/domain"
	"mi-api-go/src/videojuegos/domain/entities"
)

type MysqlVideojuego struct {
	conn *sql.DB
}

func NewMysqlVideojuegoRepository() domain.IVideojuego {
	conn := core.GetDB()
	return &MysqlVideojuego{conn: conn}
}

func (mysql *MysqlVideojuego) Save(videojuego entities.Videojuego) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO videojuegos (nombre, descripcion, genero, plataforma, precio, deleted) VALUES (?, ?, ?, ?, ?, ?)",
		videojuego.Nombre,
		videojuego.Descripcion,
		videojuego.Genero,
		videojuego.Plataforma,
		videojuego.Precio,
		videojuego.Deleted,
	)
	if err != nil {
		log.Println("Error al guardar el videojuego:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	videojuego.SetID(int(idInserted))
	return nil
}

func (mysql *MysqlVideojuego) Update(id int, videojuego entities.Videojuego) error {
	result, err := mysql.conn.Exec(
		"UPDATE videojuegos SET nombre = ?, descripcion = ?, genero = ?, plataforma = ?, precio = ?, deleted = ? WHERE id = ?",
		videojuego.Nombre,
		videojuego.Descripcion,
		videojuego.Genero,
		videojuego.Plataforma,
		videojuego.Precio,
		videojuego.Deleted,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el videojuego:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el videojuego con ID:", id)
		return fmt.Errorf("videojuego con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlVideojuego) Delete(id int) error {
	_, err := mysql.conn.Exec("UPDATE videojuegos SET deleted = true WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar (soft delete) el videojuego:", err)
		return err
	}
	return nil
}

func (mysql *MysqlVideojuego) FindByID(id int) (entities.Videojuego, error) {
	var videojuego entities.Videojuego
	row := mysql.conn.QueryRow("SELECT id, nombre, descripcion, genero, plataforma, precio, deleted FROM videojuegos WHERE id = ?", id)

	err := row.Scan(
		&videojuego.ID,
		&videojuego.Nombre,
		&videojuego.Descripcion,
		&videojuego.Genero,
		&videojuego.Plataforma,
		&videojuego.Precio,
		&videojuego.Deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Videojuego no encontrado:", err)
			return entities.Videojuego{}, fmt.Errorf("videojuego con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el videojuego por ID:", err)
		return entities.Videojuego{}, err
	}

	return videojuego, nil
}

func (mysql *MysqlVideojuego) GetAll() ([]entities.Videojuego, error) {
	var videojuegos []entities.Videojuego

	rows, err := mysql.conn.Query("SELECT id, nombre, descripcion, genero, plataforma, precio, deleted FROM videojuegos")
	if err != nil {
		log.Println("Error al obtener todos los videojuegos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var videojuego entities.Videojuego
		err := rows.Scan(
			&videojuego.ID,
			&videojuego.Nombre,
			&videojuego.Descripcion,
			&videojuego.Genero,
			&videojuego.Plataforma,
			&videojuego.Precio,
			&videojuego.Deleted,
		)
		if err != nil {
			log.Println("Error al escanear videojuego:", err)
			return nil, err
		}
		videojuegos = append(videojuegos, videojuego)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return videojuegos, nil
}

func (mysql *MysqlVideojuego) GetByStatus(deleted bool) ([]entities.Videojuego, error) {
	var videojuegos []entities.Videojuego

	rows, err := mysql.conn.Query("SELECT id, nombre, descripcion, genero, plataforma, precio, deleted FROM videojuegos WHERE deleted = ?", deleted)
	if err != nil {
		log.Println("Error al obtener videojuegos por estado:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var videojuego entities.Videojuego
		err := rows.Scan(
			&videojuego.ID,
			&videojuego.Nombre,
			&videojuego.Descripcion,
			&videojuego.Genero,
			&videojuego.Plataforma,
			&videojuego.Precio,
			&videojuego.Deleted,
		)
		if err != nil {
			log.Println("Error al filtrar los videojuegos:", err)
			return nil, err
		}
		videojuegos = append(videojuegos, videojuego)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al filtrar los videojuegos:", err)
		return nil, err
	}

	return videojuegos, nil
}
