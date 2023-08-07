package main

import (
	"fmt"
	"log"
	"os"

	"application/datamodels"
	"application/interceptor"
	"application/persistantlayer"

	"github.com/gin-gonic/gin"
)

type mongoDbController struct {
	mbController persistantlayer.MongoDbInterface
}

func (structController mongoDbController) Add (row datamodels.Row) {
	err := structController.mbController.CreateRow(row)
	if err != nil {
		fmt.Println("errors : ", err)
		return
	}
	fmt.Println("Row Added!")
}

func (structController mongoDbController) GetAllRows () {
	rows, err := structController.mbController.GetAll()
	if err != nil {
		fmt.Println("errors : ", err)
		return
	}
	fmt.Println(rows)
}

func (structController mongoDbController) Update (id int, updatedRow datamodels.Row) {
	err := structController.mbController.UpdateRow(id, updatedRow)
	if err != nil {
		fmt.Println("errors : ", err)
		return
	}
	fmt.Println("Updated the row")
}

func (structController mongoDbController) Delete (id int) {
	err := structController.mbController.DeleteRow(id)
	if err != nil {
		fmt.Println("errors : ", err)
		return 
	}
	fmt.Println("Row was Deleted")
}


type GinController struct {
	ginControllerHandler persistantlayer.GinInterfaceFunction
}

func (ginControllerObject GinController) HttpPost(c *gin.Context)  {
	err := ginControllerObject.ginControllerHandler.CreateRow(c)
	fmt.Println(err)
	if err != nil {
		interceptor.SendErrRes(c, err, "Failure to create a user", 200)
		return
	}
	interceptor.SendSuccessRes(c, "User was created successfully", 200)
	return
}

func (ginControllerObject GinController) HttpGetAll(c *gin.Context) {
	allUser, err := ginControllerObject.ginControllerHandler.GetAll()
	if err != nil {
		interceptor.SendErrRes(c, err, "Could not find the users at this moment.", 404)
	}
	interceptor.SendSuccessRes(c, allUser, 200)
}

func (ginControllerObject GinController) HttpPatch(c *gin.Context) {
	err := ginControllerObject.ginControllerHandler.UpdateRow(c)
	if err !=  nil {
		interceptor.SendErrRes(c, err, "Could not update the user at this moment.", 500)
	}
	interceptor.SendSuccessRes(c, "Updated User", 200)
}

func (ginControllerObject GinController) HttpDelete(c *gin.Context) {
	err := ginControllerObject.ginControllerHandler.DeleteRow(c)
	if err != nil {
		interceptor.SendErrRes(c, err, "Could not delete the user at this moment.", 500)
	}
	interceptor.SendSuccessRes(c, "Deleted the user successfully", 200)
}

func main() {
	applicationController := GinController{ *persistantlayer.InitiateGinInterface() }

	server := gin.Default()

	logfile, err := os.Create("app.log")
	defer logfile.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logfile)
	log.Println("Starting the application...")

	server.POST("/user", applicationController.HttpPost)
	server.GET("/user", applicationController.HttpGetAll)
	server.PATCH("/user", applicationController.HttpPatch)
	server.DELETE("/user/:id", applicationController.HttpDelete)

	server.Run()
}
