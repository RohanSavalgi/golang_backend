package main

import (
	"application/controller"
	"application/db"
	"application/persistantlayer"
	envLoader "application/server"

	"github.com/gin-gonic/gin"
)

func main() {
	envLoader.LoadEnv()
	db.CreateConnection()
	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization()}
	server := gin.Default()

	server.POST("/user", applicationController.HttpPost)
	server.GET("/user", applicationController.HttpGetAll)
	server.PUT("/user", applicationController.HttpPatch)
	server.DELETE("/user/:id", applicationController.HttpDelete)

	server.Run()
}