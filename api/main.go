package main

import (
	"application/cmd"
	_ "application/datamodels"
	"application/db"
	_ "application/docs"
	"application/logger"
	"application/mq"
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
// @SecurityDefinitions BearerToken
// @host localhost:8080
// @BasePath /application
func main() {
	mainServer := server.InitServer()
	resty.CreateRestyClient()

	cmd.ServerRoutesSetupUp(mainServer)

	if orgId, err :=  mq.CreateNewUser("nascar", "nascar@gmail.com"); err != nil {
		logger.ThrowErrorLog(err)
	} else {
		logger.ThrowDebugLog(orgId)
	}

	go pubsub.RecieveMessageFromGoRoutine()

	server.Listen(mainServer)
	
}