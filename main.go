package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"mi-api-go/src/core"
	usuariosRut "mi-api-go/src/usuarios/infrastructure/routes_usuarios"
	videojuegoRut "mi-api-go/src/videojuegos/infrastructure/routes_videojuegos"
)

func main() {

	core.InitDB()

	r := gin.Default()

	usuariosRouter := usuariosRut.NewRouter(r)
	usuariosRouter.Run()

	juegosRouter := videojuegoRut.NewRouter(r)
	juegosRouter.Run()

	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
