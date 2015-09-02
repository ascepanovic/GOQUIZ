package controllers

import(
  "github.com/gin-gonic/gin"
)

//HomeHandler this function is using ginContext
func Home(c *gin.Context) {
	//and sand some response do
	c.JSON(200, gin.H{"message": "Hello from DMAS Go Quiz!"})
}
