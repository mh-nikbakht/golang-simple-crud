// config/routes.go
package config

import (
	"golang_learning/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, userController controllers.UserController) {
	basepath := server.Group("/v1")
	userController.RegisterRoutes(basepath)

}
