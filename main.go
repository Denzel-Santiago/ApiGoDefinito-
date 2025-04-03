package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mi-api-go/src/core"
	usuariosRut "mi-api-go/src/usuarios/infrastructure/routes_usuarios"
	videojuegoRut "mi-api-go/src/videojuegos/infrastructure/routes_videojuegos"
)

func main() {
	// Prueba temporal de bcrypt (eliminar después de verificar)
	testPasswordHashing()

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

// Función temporal para probar el hashing (eliminar después)
func testPasswordHashing() {
	fmt.Println("\n=== PRUEBA TEMPORAL DE HASHING ===")

	password := "miContraseñaSegura123"

	// 1. Generar hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error generando hash:", err)
	}

	fmt.Println("Contraseña original:", password)
	fmt.Println("Hash generado:", string(hashedPassword))

	// 2. Verificar que el hash funciona
	fmt.Println("\nVerificando contraseña correcta:")
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == nil {
		fmt.Println("✅ La contraseña coincide!")
	} else {
		fmt.Println("❌ Error en verificación:", err)
	}

	// 3. Verificar con contraseña incorrecta
	fmt.Println("\nVerificando contraseña incorrecta:")
	wrongPassword := "contraseñaEquivocada"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(wrongPassword))
	if err != nil {
		fmt.Println("✅ Correcto: la contraseña no coincide (error esperado)")
		fmt.Println("Error devuelto:", err)
	} else {
		fmt.Println("❌ Peligro: la contraseña incorrecta fue aceptada!")
	}

	fmt.Println("=== FIN DE PRUEBA TEMPORAL ===\n")

	// Opcional: pausa para ver los resultados en consola
	fmt.Print("Presiona Enter para continuar con el inicio del servidor...")
	fmt.Scanln()
}
