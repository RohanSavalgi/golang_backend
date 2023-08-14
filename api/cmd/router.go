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

	mainRouter := router.Group("/application") 
	{
		mainRouter.POST("/user", authMiddleware, middleware.VerifyContentType(middleware.ContentTypeJSON), applicationController.HttpPost)
		mainRouter.GET("/user", authMiddleware, applicationController.HttpGetAll)
		mainRouter.PUT("/user", authMiddleware, middleware.VerifyContentType(middleware.ContentTypeJSON),applicationController.HttpPatch)
		mainRouter.DELETE("/user/:id", authMiddleware, applicationController.HttpDelete)
	}
}