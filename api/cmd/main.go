package main

import (
	"application/db"
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

	server.Listen(mainServer)
}