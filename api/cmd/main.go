package main

import (
	"encoding/json"
	"net/http"

	"application/db"
	"application/logger"
	"application/mq/pubsub"
	"application/resty"
	"application/server"
	envLoader "application/server"
)

func init() {
	envLoader.LoadEnv()
	db.CreateConnection()
}

type responseFromHttp struct {
	data string `json:"data"`
	success bool `json:"success"`
	error string `json:"errror"`
	message string `json:"message"`
}

func main() {
	mainServer := server.InitServer()
	resty.CreateRestyClient()

	serverRoutesSetupUp(mainServer)

	// go pubsub.PublishMessageFromGoRoutine("tommy")

	go pubsub.RecieveMessageFromGoRoutine()

	go checkApiFromResty()

	// go pubsub.Recorder()
	server.Listen(mainServer)
	
}

func checkApiFromResty() {
	rs := resty.GetRestyClient()
	restyRes, err := rs.Send(
		rs.GetClient().R().SetAuthToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlVaalBNT0FBaHEtcmxwbnU2SDZSRyJ9.eyJpc3MiOiJodHRwczovL2Rldi1rMXFvdHltYW0xdXN0OHphLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJrVFdxU0V3c0JOVmdkdmgyQkpteHpUV0FDcTgwZERHNkBjbGllbnRzIiwiYXVkIjoiaHR0cDovLzEyNy4wLjAuMTo4MDgwL2FwcGxpY2F0aW9uL3VzZXIiLCJpYXQiOjE2OTIzNDM1MTUsImV4cCI6MTY5MjQyOTkxNSwiYXpwIjoia1RXcVNFd3NCTlZnZHZoMkJKbXh6VFdBQ3E4MGRERzYiLCJzY29wZSI6InJlYWQ6dXNlciB1cGRhdGU6dXNlciBjcmVhdGU6dXNlciIsImd0eSI6ImNsaWVudC1jcmVkZW50aWFscyIsInBlcm1pc3Npb25zIjpbInJlYWQ6dXNlciIsInVwZGF0ZTp1c2VyIiwiY3JlYXRlOnVzZXIiXX0.iVL-t6zdGT3w-b9Uq2GrQZE41gy5hKhvqr58ownyivSLc-4ExR2MbUPoBFfxD4BYzrxKwoGbgQzFb4-Z77zgc2XcCXL9jBMbxcEVZBUeum0EYPfbxgSZmRGHJl6nJ_ynH-gi1ZcWmTu1tQIoTTikP5M2-ppziADN3jHldeEmwz9awrn6ehjs4NICYLZkw0SVmHuJApVAa1Lw_3aThBx7egJgmoeYuAmFghumFOwlN9B4PTEXuTY2wcpW9Q4wIk6N2zZCZG3Vj9ECDSUg_G-92uX1EKgwZoBxMES14YH-uMN057ptUnCbnC0G8QDA-BaYZq2tWNkadOBJ9prbClXNYg"),
		"http://127.0.0.1:8080/application/user",
		resty.GET,
	)

	checkResponseOutput, err := rs.CheckResponse(restyRes, err, http.StatusOK, resty.USR_MGMT)
	if err != nil {
		logger.ThrowErrorLog("Error in checking response.")
	}
	var newObj responseFromHttp
	err = json.Unmarshal(checkResponseOutput, newObj)
	logger.ThrowDebugLog(newObj.data)
}