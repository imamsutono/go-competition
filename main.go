package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kompit-recruitment/backend/generated/api"
	"github.com/kompit-recruitment/backend/handler"
	"github.com/kompit-recruitment/backend/initializers"
)

func main() {
	router := gin.Default()

	var h api.ServerInterface = handler.New()

	initializers.LoadConfig("app.env")
	initializers.InitDatabase()

	//TODO: You can add your own middlewares here
	router.Use(gin.Recovery())
	// router.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}))

	api.RegisterHandlers(router, h)

	router.Run(":8080")
}
