package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
)

type Opportunity struct {
	Id          string `json:"identifier_opportunity"`
	Name        string `json:"name_opportunity"`
	Description string `json:"description_of_opportunity"`
}

type Opportunities struct {
	ListOfOpportunities []Opportunity
}

func InitializeRoutes(router *gin.Engine) {
	slog.Info("Initialize routes of the server")

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/opportunities", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Opportunities": Opportunity{
					Id:          "1",
					Name:        "Golang Senior Developer",
					Description: "Golang Senior Developer, Remote Opportunity of Brazil",
				},
			})
		})
	}

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"health": "check",
		})
	})

}
