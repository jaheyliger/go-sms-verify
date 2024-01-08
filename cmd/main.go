package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaheyliger/go-sms-verify/api"
)

func main() {
	
	router := gin.Default()

	//initializa config
	app := api.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8000")
}