package middleware

import (
	"campaigns-restapi/auth"
	"campaigns-restapi/helper"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(authService auth.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.GetHeader("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			errorResponse := helper.ApiResponse(false, "Error occured", "Invalid of authorization value")
			context.JSON(http.StatusUnauthorized, errorResponse)
			return
		}

		var tokenStr string
		arrayAuthorizationValue := strings.Split(authorizationHeader, " ")
		if len(arrayAuthorizationValue) == 2 {
			tokenStr = arrayAuthorizationValue[1]
		}

		token, err := authService.ValidateToken(tokenStr)
		if err != nil {
			errorResponse := helper.ApiResponse(false, "Error occured", "Failed to validate token")
			context.JSON(http.StatusUnauthorized, errorResponse)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			errorResponse := helper.ApiResponse(false, "Error occured", "Failed to claims jwt token")
			context.JSON(http.StatusUnauthorized, errorResponse)
			return
		}

		userId := int(claim["userId"].(float64))

		user, err := authService.FindUserById(userId)
		if err != nil {
			errorResponse := helper.ApiResponse(false, "Error occured", "Unauthorized")
			context.JSON(http.StatusUnauthorized, errorResponse)
			return
		}

		context.Set("user", user)
	}
}