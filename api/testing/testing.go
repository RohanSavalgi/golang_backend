package testing

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	POST_REQUEST_BODY = `{
		"ID": 1,
		"Name": "test_name",
		"Email": "test_name@test_name.com"
	}`
)

func GetRequest(mainTestUri string, testUri string, testMethod gin.HandlerFunc) (*httptest.ResponseRecorder, error) {
	router := gin.Default()

	router.GET(testUri, testMethod)

	recorder := httptest.NewRecorder()

	testRequest, testRequestError := http.NewRequest("GET", testUri, nil)
	if testRequestError != nil {
		return nil, testRequestError
	}
	router.ServeHTTP(recorder, testRequest)
	return recorder, nil
}

func PostRequest(mainTestUri string, testUri string, testMethod gin.HandlerFunc, requestBody string) (*httptest.ResponseRecorder, error) {
	router := gin.Default()

	router.POST(testUri, testMethod)

	recorder := httptest.NewRecorder()

	testRequest, testRequestError := http.NewRequest("POST", testUri, strings.NewReader(requestBody))
	if testRequestError != nil {
		return nil, testRequestError
	}
	router.ServeHTTP(recorder, testRequest)
	return recorder, nil
}

func PatchRequest(mainTestUri string, testUri string, testMethod gin.HandlerFunc, requestBody string) (*httptest.ResponseRecorder, error) {
	router := gin.Default()

	router.PUT(testUri, testMethod)

	recorder := httptest.NewRecorder()

	testRequest, testRequestError := http.NewRequest("PUT", testUri, strings.NewReader(requestBody))
	if testRequestError != nil {
		return nil, testRequestError
	}
	router.ServeHTTP(recorder, testRequest)
	return recorder, nil
}

func DeleteRequest(mainTestUri string, testUri string, testMethod gin.HandlerFunc) (*httptest.ResponseRecorder, error) {
	router := gin.Default()

	recorder := httptest.NewRecorder()

	testRequest, testRequestError := http.NewRequest("DELETE", testUri, nil)
	if testRequestError != nil {
		return nil, testRequestError
	}
	router.ServeHTTP(recorder, testRequest)
	return recorder, nil
}