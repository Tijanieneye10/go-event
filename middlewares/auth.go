package middlewares

import (
	"github.com/Tijanieneye10/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(context *gin.Context) {
	//get token
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized user"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
