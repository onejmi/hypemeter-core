package main

import (
	"github.com/heroku/hypemeter-core/lib/component"
	"github.com/heroku/hypemeter-core/lib/data"
	"github.com/heroku/hypemeter-core/lib/util"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8888"
	}

	//connect to DB
	data.Connect()

	//run cronjobs
	util.SetupCronJobs()

	router := gin.Default()

	router.POST("/api/login", component.HandleLogin)

	router.Run(":" + port)
}
