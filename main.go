package main

import (
	"github.com/heroku/hypemeter-core/lib/component"
	"github.com/heroku/hypemeter-core/lib/data"
	"github.com/heroku/hypemeter-core/lib/util"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

const apiBasePath = "api/v1"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8888"
	}

	//connect to DB
	data.Connect()

	//cron jobs
	util.SetupCronJobs()

	router := gin.Default()
	//groups
	authorized := router.Group(apiBasePath, data.Authorize)
	authorized.GET("/test", data.Test)
	authorized.GET("/profile/view", component.ViewProfile)

	router.POST("/login", component.HandleLogin)

	router.Run(":" + port)
}
