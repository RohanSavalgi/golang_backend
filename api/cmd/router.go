package main

import (
	"application/auth"
	"application/controller"
	"application/logger"
	"application/middleware"
	"application/persistantlayer"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

func serverRoutesSetupUp(router *gin.Engine) {
	logger.ThrowDebugLog("Setting up the routes for the server.")

	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization()}
	authMiddleware := adapter.Wrap(auth.AuthenticationMiddleware())
	/*
	Owner - Highest - Create, Read, Update
	Admin - Middle - Read, Update
	Employee - Lowest - Read

	*/
	mainRouter := router.Group("/application") 
	{
		mainRouter.GET("/user", authMiddleware, auth.CheckPermission([]string{"read:user"}, true), applicationController.HttpGetAll)
		mainRouter.POST("/user", authMiddleware, auth.CheckPermission([]string{"create:user"}, true), middleware.VerifyContentType(middleware.ContentTypeJSON), applicationController.HttpPost)
		mainRouter.PUT("/user", authMiddleware, auth.CheckPermission([]string{"update:user"}, true), middleware.VerifyContentType(middleware.ContentTypeJSON),applicationController.HttpPatch)
		mainRouter.DELETE("/user/:id", authMiddleware, applicationController.HttpDelete)
	}
}