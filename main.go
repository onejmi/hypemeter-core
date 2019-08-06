package main

import (
	"github.com/heroku/hypemeter-core/lib/component"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.POST("/login", component.HandleLogin)

	router.Run(":" + port)
}
