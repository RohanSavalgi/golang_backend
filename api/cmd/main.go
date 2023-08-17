package main

import (
	"application/db"
	"application/server"
	envLoader "application/server"
	"application/mq/pubsub"
)

func init() {
	envLoader.LoadEnv()
	db.CreateConnection()
}

func main() {
	mainServer := server.InitServer()

	serverRoutesSetupUp(mainServer)

	go pubsub.PublishMessageFromGoRoutine()

	server.Listen(mainServer)
}