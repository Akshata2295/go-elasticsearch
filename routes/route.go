package routes

import (
	"fmt"
	controllers "go-elasticsearch/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(func(c *gin.Context) {
		//allow all
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers",  "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			platform := c.GetHeader("User-agent")
			fmt.Println("platofor", platform)
		}
		c.Next()
	})
	r.POST("/documents", controllers.CreateDocuments)
	r.GET("/search", controllers.SearchDocument)
	return r
}