package main

import (
	"fmt"
	"log"
	"time"

	"mi-api-go/src/core"
	usuariosRut "mi-api-go/src/usuarios/infrastructure/routes_usuarios"
	videojuegoRut "mi-api-go/src/videojuegos/infrastructure/routes_videojuegos"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Prueba temporal de hashing (puedes eliminar esta sección después de verificar)
	testPasswordHashing()

	// Inicializar la conexión a la base de datos
	core.InitDB()

	// Crear el router de Gin
	r := gin.Default()

	// Configuración detallada de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todos los orígenes (en producción, especifica tus dominios)
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar rutas de usuarios
	usuariosRouter := usuariosRut.NewRouter(r)
	usuariosRouter.Run()

	// Configurar rutas de videojuegos
	juegosRouter := videojuegoRut.NewRouter(r)
	juegosRouter.Run()

	// Iniciar el servidor en el puerto 8000
	fmt.Println("Servidor iniciado en http://localhost:8000")
	err := r.Run(":8000")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

// testPasswordHashing es una función temporal para probar el hashing de contraseñas
// Puede eliminarse después de verificar que el hashing funciona correctamente
func testPasswordHashing() {
	fmt.Println("\n=== PRUEBA TEMPORAL DE HASHING ===")

	password := "miContraseñaSegura123"

	// 1. Generar hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error generando hash:", err)
	}

	fmt.Println("Contraseña original:", password)
	fmt.Println("Hash generado:", string(hashedPassword))

	// 2. Verificar contraseña correcta
	fmt.Println("\nVerificando contraseña correcta:")
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == nil {
		fmt.Println("✅ La contraseña coincide!")
	} else {
		fmt.Println("❌ Error en verificación:", err)
	}

	// 3. Verificar contraseña incorrecta
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
}
