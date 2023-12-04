package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run(":8080")

}
