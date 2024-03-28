package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/David-Carr-C/Ecommerce/controllers"
	"gitlab.com/David-Carr-C/Ecommerce/services"
	"gorm.io/gorm"
)

func UserEndpoint(api *gin.RouterGroup, dataBase *gorm.DB) {
	api.POST("/login", controllers.LoginHandler(services.Login, dataBase))
}
