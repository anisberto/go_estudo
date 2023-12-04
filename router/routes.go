package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRoutes(router *gin.Engine) {

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"health": "checked",
		})
	})

	router.GET("/product", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id":          1,
			"opportunity": "Senior Golang Developer",
		})
	})
}
