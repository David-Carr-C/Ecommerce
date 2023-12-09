package main

import (
	"gitlab.com/David-Carr-C/Ecommerce/configs"
	"gitlab.com/David-Carr-C/Ecommerce/routers"
)

func main() {
	// Se obtiene la base de datos
	dataBase := configs.SetupDatabase()

	// Se cargan las rutas
	router := routers.SetupRouter(dataBase)

	// Se inicia el servidor en el puerto 8080
	router.Run(":8080")
}
