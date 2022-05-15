package controller

import (
	"net/http"

	"github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/service"
	"github.com/gin-gonic/gin"
)


type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *authController {
     return &authController{authService}
}

func (auth *authController) Login(ctx *gin.Context) {

}

func (auth *authController) Register(ctx *gin.Context) {
		var user model.UserRequest
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}	
		createdUser , err := auth.authService.RegisterUser(user)
		if err != nil {
			response := gin.H{
				"error": err.Error(),
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := gin.H{
			"message": "user registerd successfully",
			"user": createdUser,
		}
		ctx.JSON(http.StatusCreated, response)
	

}