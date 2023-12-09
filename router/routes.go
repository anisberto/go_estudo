package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
)

func InitializeRoutes(router *gin.Engine) {
	slog.Info("Initialize routes of the server")
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"health": "check",
		})
	})

	router.GET("/opportunities", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"id":          1,
			"opportunity": "Senior Golang Developer",
		})
	})
}
