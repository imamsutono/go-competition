package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kompit-recruitment/backend/generated/api"
	"github.com/kompit-recruitment/backend/handler"
)

func main() {
	router := gin.Default()

	//TODO: You may load the configurations (e.g. database access) or other
	// 		objects here and pass them to the handler.

	// You can use database using environment variable `DATABASE_URL`.

	var h api.ServerInterface = handler.New()

	//TODO: You can add your own middlewares here
	router.Use(gin.Recovery())

	api.RegisterHandlers(router, h)

	router.Run(":8080")
}
