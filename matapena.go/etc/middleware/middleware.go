package middleware

import (
	"app/etc/base"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func resError(c *gin.Context, msg string, code int) {
	c.JSON(200, base.JSONRes{
		Success: false,
		Data:    nil,
		Error:   base.JSONError{Message: msg, Code: code},
	})
	return
}

// TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := viper.GetString("reqToken")

	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-ReqToken")

		if token == "" {
			resError(c, "API token X-ReqToken", 401)
			c.Abort()
			return
		}

		if token != requiredToken {
			resError(c, "Invalid API token", 401)
			c.Abort()
			return
		}

		c.Next()
	}
}
