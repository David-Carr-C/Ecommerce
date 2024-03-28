package routers

import (
	"net/http"

	"gitlab.com/David-Carr-C/Ecommerce/controllers"
	"gitlab.com/David-Carr-C/Ecommerce/services"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(dataBase *gorm.DB) *gin.Engine {
	// Se crea el framework Gin
	router := gin.Default()

	// Se permite CORS
	router.Use(corsMiddleware())

	// Se define el frontend
	router.Use(static.Serve("/", static.LocalFile("./templates", true)))

	// Se define el grupo de rutas
	api := router.Group("/api")

	// Se define la ruta con su respectivo controlador y servicio
	api.POST("/create-user", controllers.CreateUserHandler(services.CreateUser, dataBase))
	api.GET("/user/:id", controllers.GetUserHandler(services.GetUserByID, dataBase))
	api.PUT("/user/:id", controllers.UpdateUserHandler(services.UpdateUser, dataBase))
	api.DELETE("/user/:id", controllers.DeleteUserHandler(services.DeleteUser, dataBase))

	// Se definen rutas en archivos separados
	UserEndpoint(api, dataBase)
	SendRouter(router)

	// 404 - React
	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./templates")
	})

	// Retorna Gin
	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Solo puede haber un dominio permitido por Allow-Credentials
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // es necesario para las cookies

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
