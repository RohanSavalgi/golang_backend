package main

import (
	"application/db"
	"application/mq/pubsub"
	"application/server"
	envLoader "application/server"
)

func init() {
	envLoader.LoadEnv()
	db.CreateConnection()
}

func main() {
	mainServer := server.InitServer()

	serverRoutesSetupUp(mainServer)

	// go pubsub.PublishMessageFromGoRoutine("tommy")

	go pubsub.RecieveMessageFromGoRoutine()

	// go pubsub.OtherPublish()

	// go pubsub.Recorder()

	server.Listen(mainServer)
}