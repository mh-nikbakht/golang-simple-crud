// ioc/controllers.go
package ioc

import (
	"golang_learning/controllers"
)

var (
	UserController controllers.UserController
)

func InitializeControllers() {
	UserController = controllers.New(UserService)

}
