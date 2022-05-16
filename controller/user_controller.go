package controller

import (
	"github.com/aminyasser/todo-list/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	authService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) *userController {
	return &userController{userService, jwtService}
}

func (user *userController) Profile(ctx *gin.Context) {

}

func (user *userController) Update(ctx *gin.Context) {
	
}
