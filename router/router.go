package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Init() {
	router := gin.Default()
	InitializeRoutes(router)

	err := router.Run(":8080")

	if err != nil {
		slog.Error("Error during initialize server, ", err)
	}

}
