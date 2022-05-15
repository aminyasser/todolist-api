package controller

import (
	"net/http"
	"strconv"

	// "github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/entity/request"
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/service"
	"github.com/gin-gonic/gin"
)


type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService , jwtService service.JWTService) *authController {
     return &authController{authService , jwtService}
}

func (auth *authController) Login(ctx *gin.Context) {
	var userLogin request.UserLogin
	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		response := response.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	
	user , err := auth.authService.VerifyUser(userLogin.Email, userLogin.Password)
	if err != nil {
		response := response.Error( err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token := auth.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
	user.Token = token
		
	response := response.Success("user login successfully", user)
	ctx.JSON(http.StatusCreated, response)

}

func (auth *authController) Register(ctx *gin.Context) {
		var user request.UserRegister
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			response := response.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}	
		createdUser , err := auth.authService.RegisterUser(user)
		if err != nil {
			response := response.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token := auth.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		
		response := response.Success("user registerd successfully", createdUser)
		ctx.JSON(http.StatusCreated, response)
	

}