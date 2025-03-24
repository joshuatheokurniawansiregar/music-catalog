package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc{
	var secretKey string = configs.GetConfig().Service.SecretKey
	return func(c *gin.Context){
		var header string = c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("token is missing"))
			return
		}
		userId, userName, err := jwt.ValidateToken(header, secretKey)
		if err != nil{
			c.AbortWithError(http.StatusUnauthorized, errors.New("token is invalid"))	
			return
		}
		c.Set("userId",userId)
		c.Set("username", userName)
		c.Next()
	}
}