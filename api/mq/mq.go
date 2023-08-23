package mq

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	dto "application/dto/organisation"
	userDto "application/dto/user"
	"application/logger"
	"application/resty"
)

func CreateOranizationForUser(name string) (interface{}, error) {
	miniOrgData := dto.CreateOrgRequestModel{
		Name: name, 
		DisplayName: name,
	}
	rs := resty.GetRestyClient()

	organizationCheck, organizationCheckError := rs.Send(
		rs.GetClient().R().
			SetAuthToken(os.Getenv("APPLICATION_ALL_ACCESS_TOKEN")),
			"https://" + os.Getenv("APPLICATION_DOMAIN") + "/api/v2/organizations/name/" + miniOrgData.Name,
			resty.GET,
	)

	if organizationCheckError != nil {
		return nil, errors.New("Failed to check for duplicate organization name")
	}

	if organizationCheck.StatusCode() != http.StatusNotFound {
		orgCheck, err := rs.CheckResponse(organizationCheck, organizationCheckError, http.StatusOK, resty.AUTH0)
		if err != nil {
			logger.ThrowErrorLog("Error checking organization data in auth0")
			return nil, err
		}

		result := dto.CreateOrgResponseModel{}
		if err := json.Unmarshal(orgCheck, &result); err != nil {
			logger.ThrowErrorLog("Error in unmarshalling data to check for duplicate organization")
			return nil, err
		}

		logger.ThrowErrorLog("Organization with this name is already found! Contining with the next process")
		return result, nil
	}

	restyRequest, restyRequestErr := rs.Send(
		rs.GetClient().R().
			SetAuthToken(os.Getenv("APPLICATION_ALL_ACCESS_TOKEN")).
			SetBody(miniOrgData),
			"https://" + os.Getenv("APPLICATION_DOMAIN") + "/api/v2/organizations",
			resty.POST,
		)

	data, err := rs.CheckResponse(restyRequest, restyRequestErr, http.StatusCreated, resty.AUTH0)
	if err != nil {
		logger.ThrowErrorLog(err)
	}


	result := dto.CreateOrgResponseModel{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateNewUser(name string, email string) (interface{}, error) {
	rs := resty.GetRestyClient()

	requestBody := userDto.CreateUserRequestModel{
		Name: name,
		Email: email,
		Connection: "",
		Password: "random",
	}
	userCreationReq, err := rs.Send(
		rs.GetClient().R().
			SetAuthToken(os.Getenv("APPLICATION_ALL_ACCESS_TOKEN")).
			SetBody(requestBody),
		"https://" + os.Getenv("APPLICATION_DOMAIN") + "/api/v2/user",
		resty.POST,
	)

	responseInBytes, err := rs.CheckResponse(userCreationReq, err, http.StatusCreated, resty.AUTH0)
	if err != nil {
		return nil, err
	}

	userCreationResponse := userDto.CreateUserRequestModel{}
	if err := json.Unmarshal(responseInBytes, &userCreationResponse); err != nil {
		logger.ThrowErrorLog("Error in unmarshalling data!")
		return nil, err
	}

	return userCreationResponse, nil
}