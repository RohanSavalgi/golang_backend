package unit_tests

import (
	"net/http"
	"testing"

	"application/controller"
	"application/db"
	"application/logger"
	"application/persistantlayer"
	applicationTesting "application/testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	testRouter = gin.Default()
)

var (
	GET_TEST_BODY = `{
		"data": [
		  {
			"ID": 1,
			"Name": "rohan is the best",
			"Email": "rohan@rohan.com"
		  },
		  {
			"ID": 2,
			"Name": "panda",
			"Email": "panda@betsol.com"
		  }
		],
		"success": true,
		"error": null,
		"message": ""
	  }`
	POST_REQUEST_BODY = `{
			"ID": 3,
			"Name": "panda",
			"Email": "panda@betsol.com"
	}`
	POST_TEST_BODY = `{
			"data": "User was created successfully",
			"success": true,
			"error": null,
			"message": ""
	}`
	PUT_TEST_BODY = `{
		"data": "Updated User",
		"success": true,
		"error": null,
		"message": ""
	  }`
)

func TestHttpGetAll(t *testing.T) {
	db.CreateConnection()
	applicationController := controller.PgDbController{ PgControllerHandler : *persistantlayer.PostgresInitilization() }

	testRecorder, testErrors := applicationTesting.GetRequest("/test-user","/test-user",applicationController.HttpGetAll)
	if testErrors != nil {
		logger.ThrowErrorLog(testErrors)
		return
	}

	assert.Equal(t, http.StatusOK, testRecorder.Code)
	assert.JSONEq(t, GET_TEST_BODY, testRecorder.Body.String())
}

func TestHttpPost(t *testing.T) {
	db.CreateConnection()
	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization() }
	
	testRecorder, testErrors := applicationTesting.PostRequest("/test-user", "/test-user", applicationController.HttpPost, POST_REQUEST_BODY)
	if testErrors != nil {
		logger.ThrowErrorLog(testErrors)
		return 
	}
	assert.Equal(t, http.StatusOK, testRecorder.Code)
	assert.JSONEq(t, POST_TEST_BODY, testRecorder.Body.String())
}

func TestHttpPut(t *testing.T) {
	db.CreateConnection()
	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization() }
	
	testRecorder, testErrors := applicationTesting.PatchRequest("/test-user", "/test-user", applicationController.HttpPatch , POST_REQUEST_BODY)
	if testErrors != nil {
		logger.ThrowErrorLog(testErrors)
		return 
	}
	assert.Equal(t, http.StatusOK, testRecorder.Code)
	assert.JSONEq(t, PUT_TEST_BODY, testRecorder.Body.String())
}

func TestHttpDelete(t *testing.T) {
	db.CreateConnection()
	applicationController := controller.PgDbController{ PgControllerHandler: *persistantlayer.PostgresInitilization() }
	
	testRecorder, testErrors := applicationTesting.PatchRequest("/test-user", "/test-user", applicationController.HttpPatch , POST_REQUEST_BODY)
	if testErrors != nil {
		logger.ThrowErrorLog(testErrors)
		return 
	}
	assert.Equal(t, http.StatusOK, testRecorder.Code)
	assert.JSONEq(t, PUT_TEST_BODY, testRecorder.Body.String())
}