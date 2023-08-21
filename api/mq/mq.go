package mq

import (
	"strings"

	"application/logger"
	"application/resty"
)

var(
	URI = "https://" + "dev-k1qotymam1ust8za.us.auth0.com" + "/api/v2/organizations"
)

func CreateNewUserInAuth0() {
	rs := resty.GetRestyClient()
	payload := strings.NewReader(`{
		"name": "new_rohan_org",
		"display_name": "new_rohan_orgnew_rohan_org",
		"branding": [
			{
				"logo_url": "",
				"colors": [
					{
						"primary": "",
						"page_background": ""
					}
				]
			}
		],
		"metadata": [
			{}
		],
		"enabled_connections": [
		]
	}`)

	_, err := rs.Send(
		rs.GetClient().R().SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlVaalBNT0FBaHEtcmxwbnU2SDZSRyJ9.eyJpc3MiOiJodHRwczovL2Rldi1rMXFvdHltYW0xdXN0OHphLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJrVFdxU0V3c0JOVmdkdmgyQkpteHpUV0FDcTgwZERHNkBjbGllbnRzIiwiYXVkIjoiaHR0cDovLzEyNy4wLjAuMTo4MDgwL2FwcGxpY2F0aW9uL3VzZXIiLCJpYXQiOjE2OTI1ODczNDksImV4cCI6MTY5MjY3Mzc0OSwiYXpwIjoia1RXcVNFd3NCTlZnZHZoMkJKbXh6VFdBQ3E4MGRERzYiLCJzY29wZSI6InJlYWQ6dXNlciB1cGRhdGU6dXNlciBjcmVhdGU6dXNlciIsImd0eSI6ImNsaWVudC1jcmVkZW50aWFscyIsInBlcm1pc3Npb25zIjpbInJlYWQ6dXNlciIsInVwZGF0ZTp1c2VyIiwiY3JlYXRlOnVzZXIiXX0.QTcs65sXEHHv3mfXS8QH1tmKlNIlEF28EfGWLQi9-o8eMgZ8MqPBBF3b_ilzoR9cFjcBn55d9Y6hm5iBHqC8itLwFFja3kPCpXQyYx-fV_tYoWDyOcdOw1ddFC7YSlbH2EBFdeGmJn00JPz8qsgyE_1Guft5-IxgHhSfpLnsJPXTvnRxT3zQFoDWzHEHIZ6-nKQopIOo5AOwV-VJu1GT0L0KUwSbE6ZWeeKTER_9jIJouRSWZO8nmxITKSy8kogHOtYmhbDfhgtbD8iflsYZfzqMKDVCloXUTKCxWS1MKs8tbA3T2Y5A92oZpLOkPhxRi6ah7O686gfUcs1lEbHkuw").
		SetHeader("Cache-Control", "no-cache").
		SetBody(payload),
		URI,
		resty.POST,

	)
	if err != nil {
		logger.ThrowErrorLog(err)
		return
	}

	logger.ThrowDebugLog("Organization was created successfully!")
	return 
}