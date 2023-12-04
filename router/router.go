package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	InitializeRoutes(router)
	err := router.Run(":8080")

	if err != nil {
		return
	}

}
