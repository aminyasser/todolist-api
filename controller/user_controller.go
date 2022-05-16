package controller

import (
	"fmt"

	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) *userController {
	return &userController{userService, jwtService}
}

func (c *userController) getIdFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (user *userController) Profile(ctx *gin.Context) {
	id := user.getIdFromHeader(ctx)
	profile, err := user.userService.FindUser(id)
	if err != nil {
		response.Error(err.Error())
	} else {
		response := response.Success("OK" , profile)
		ctx.JSON(200 , response)
	}
}

func (user *userController) Update(ctx *gin.Context) {
	
}
