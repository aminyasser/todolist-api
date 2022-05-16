package middleware

import (
	"log"
	"net/http"

	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWT validates the token user 
func JWT(jwtService service.JWTService) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := response.Error("No token, Failed to process request")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token := jwtService.ValidateToken(authHeader, c)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			response := response.Error("Your token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
