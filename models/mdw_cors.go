package dmas

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware will inject different headers into our gin response writer
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("CORS middleware loaded...")

		//var url = c.Request.URL

		c.Writer.Header().Set("Access-Control-Allow-Origin", "null")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, WS, WSS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
