package main

import (
	_ "application/datamodels"
	_ "application/docs"
	"application/db"
	"application/cmd"
	"application/mq/pubsub"
	"application/resty"
	"application/server"
	envLoader "application/server"

)

func init() {
	envLoader.LoadEnv()
	db.CreateConnection()
}

// @title User Application APIs
// @version 1.0
// @description Contains all the apis for user (postgres database).
// @securityDefinitions.apiKey JWT
// @in header
// @name Bearer
// @host localhost:8080
// @BasePath /application
func main() {
	mainServer := server.InitServer()
	resty.CreateRestyClient()

	cmd.ServerRoutesSetupUp(mainServer)

	go pubsub.RecieveMessageFromGoRoutine()

	server.Listen(mainServer)
	
}