// ioc/services.go
package ioc

import (
	"golang_learning/config"
	"golang_learning/services/user"
)

var (
	UserService user.UserService
)

func InitializeServices() {
	UserService = user.NewUserService(config.UserColl, config.Ctx)
}
