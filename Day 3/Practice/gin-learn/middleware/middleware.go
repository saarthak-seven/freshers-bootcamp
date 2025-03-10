package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// // Real World middleware implementation example
// func Authenticate() gin.HandlerFunc {
// 	// Custom Logic for running before our Middleware
// 	return func(ctx *gin.Context) {
// 		if !(ctx.Request.Header.Get("Token") == "auth") {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"Message": "Token Not present",
// 			})
// 			return
// 		}
// 		ctx.Next()
// 	}

// }

// Normal Middleware
func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Message": "Token Not present",
		})
		return
	}
	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}
