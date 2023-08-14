package main

import (
	"application/controller"
	"application/persistantlayer"
	"application/logger"
	"application/middleware"

	"github.com/gin-gonic/gin"
)

func serverRoutesSetupUp(router *gin.Engine) {
	logger.ThrowDebugLog("Setting up the routes for the server.")

	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization()}

	mainRouter := router.Group("/application") 
	{
		mainRouter.POST("/user", middleware.VerifyContentType(middleware.ContentTypeJSON), applicationController.HttpPost)
		mainRouter.GET("/user", applicationController.HttpGetAll)
		mainRouter.PUT("/user", middleware.VerifyContentType(middleware.ContentTypeJSON),applicationController.HttpPatch)
		mainRouter.DELETE("/user/:id", applicationController.HttpDelete)
	}
}