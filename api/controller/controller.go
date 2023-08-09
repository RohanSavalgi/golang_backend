package main

import (
	"fmt"

	"application/db"
	"application/interceptor"
	"application/logger"
	"application/persistantlayer"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type PgDbController struct {
	pgControllerHandler persistantlayer.PostgresInterface
}

func (pgControllerObject PgDbController) HttpPost(c *gin.Context)  {
	err := pgControllerObject.pgControllerHandler.CreateRow(c)
	if err != nil {
		logger.ThrowErrorLog(err)
		interceptor.SendErrRes(c, err, "Failure to create a user", 500)
		logger.ThrowErrorLog("[Waring] : User info failed to be created.")
		return
	}
	interceptor.SendSuccessRes(c, "User was created successfully", 200)
	logger.ThrowDebugLog("User info was created successfully.")
}

func (pgControllerObject PgDbController) HttpGetAll(c *gin.Context) {
	allUser, err := pgControllerObject.pgControllerHandler.GetAll()
	if err != nil {
		interceptor.SendErrRes(c, err, "Could not find the users at this moment.", 404)
		logger.ThrowErrorLog("[Waring] : All user retrival failed.")
		return
	}
	interceptor.SendSuccessRes(c, allUser, 200)
	logger.ThrowDebugLog("Get api called and all user info is passed")
}

func (pgControllerObject PgDbController) HttpPatch(c *gin.Context) {
	err := pgControllerObject.pgControllerHandler.UpdateRow(c)
	if err !=  nil {
		interceptor.SendErrRes(c, err, "Could not update the user at this moment.", 500)
		logger.ThrowErrorLog("[Waring] : User info failed to be updated.")
		return
	}
	interceptor.SendSuccessRes(c, "Updated User", 200)
	logger.ThrowDebugLog("User info has been updated.")
}

func (pgControllerObject PgDbController) HttpDelete(c *gin.Context) {
	err := pgControllerObject.pgControllerHandler.DeleteRow(c)
	if err != nil {
		interceptor.SendErrRes(c, err, "Could not delete the user at this moment.", 500)
		logger.ThrowErrorLog("[Waring] : User info failed to be deleted.")
	}
	interceptor.SendSuccessRes(c, "Deleted the user successfully", 200)
	logger.ThrowDebugLog("User was deleted.")
}


type GinController struct {
	ginControllerHandler persistantlayer.GinInterfaceFunction
}

func (ginControllerObject GinController) HttpPost(c *gin.Context)  {
	err := ginControllerObject.ginControllerHandler.CreateRow(c)
	fmt.Println(err)
	if err != nil {
		interceptor.SendErrRes(c, err, "Failure to create a user", 200)
		logger.ThrowErrorLog("[Waring] : User info failed to be created.")
		return
	}
	interceptor.SendSuccessRes(c, "User was created successfully", 200)
	logger.ThrowDebugLog("User info was created successfully.")
}

func (ginControllerObject GinController) HttpGetAll(c *gin.Context) {
	allUser, err := ginControllerObject.ginControllerHandler.GetAll()
	if err != nil {
		interceptor.SendErrRes(c, err, "Could not find the users at this moment.", 404)
		logger.ThrowErrorLog("[Waring] : All user retrival failed.")
		return
	}
	interceptor.SendSuccessRes(c, allUser, 200)
	logger.ThrowDebugLog("Get api called and all user info is passed")
}

func (ginControllerObject GinController) HttpPatch(c *gin.Context) {
	err := ginControllerObject.ginControllerHandler.UpdateRow(c)
	if err !=  nil {
		interceptor.SendErrRes(c, err, "Could not update the user at this moment.", 500)
		logger.ThrowErrorLog("[Waring] : User info failed to be updated.")
		return
	}
	interceptor.SendSuccessRes(c, "Updated User", 200)
	logger.ThrowDebugLog("User info has been updated.")
}

func (ginControllerObject GinController) HttpDelete(c *gin.Context) {
	err := ginControllerObject.ginControllerHandler.DeleteRow(c)
	if err != nil {
		interceptor.SendErrRes(c, err, "Could not delete the user at this moment.", 500)
		logger.ThrowErrorLog("[Waring] : User info failed to be deleted.")
		return
	}
	interceptor.SendSuccessRes(c, "Deleted the user successfully", 200)
	logger.ThrowDebugLog("User was deleted.")

}

func main() {
	db.CreateConnection()
	applicationController := PgDbController{ pgControllerHandler: *persistantlayer.PostgresInitilization() }
	server := gin.Default()

	server.POST("/user", applicationController.HttpPost)
	server.GET("/user", applicationController.HttpGetAll)
	server.PUT("/user", applicationController.HttpPatch)
	server.DELETE("/user/:id", applicationController.HttpDelete)

	server.Run()
}
