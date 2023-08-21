package cmd

import (
	"application/auth"
	"application/controller"
	"application/logger"
	"application/middleware"
	"application/mq/pubsub"
	"application/persistantlayer"

	adapter "github.com/gwatts/gin-adapter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Updated import path
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "application/docs"
)

func ServerRoutesSetupUp(router *gin.Engine) {
	logger.ThrowDebugLog("Setting up the routes for the server.")

	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization()}
	authMiddleware := adapter.Wrap(auth.AuthenticationMiddleware())
	/*
	Owner - Highest - Create, Read, Update
	Admin - Middle - Read, Update
	Employee - Lowest - Read

	*/

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	mainRouter := router.Group("/application") 
	{
		mainRouter.GET("/user", authMiddleware, auth.CheckPermission([]string{"read:user"}, true), pubsub.RecorderMiddleware("GET", "USER"), applicationController.HttpGetAll)
		mainRouter.POST("/user", authMiddleware, auth.CheckPermission([]string{"create:user"}, true), pubsub.RecorderMiddleware("POST", "USER"), middleware.VerifyContentType(middleware.ContentTypeJSON), applicationController.HttpPost)
		mainRouter.PUT("/user", authMiddleware, auth.CheckPermission([]string{"update:user"}, true), pubsub.RecorderMiddleware("PUT", "USER"), middleware.VerifyContentType(middleware.ContentTypeJSON),applicationController.HttpPatch)
		mainRouter.DELETE("/user/:id", authMiddleware, applicationController.HttpDelete)
	}
}