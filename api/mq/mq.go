package mq

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

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
		Connection: "Username-Password-Authentication",
		Password: "Guru!cha1",
	}

	userExistsCheck, userExistsCheckError := rs.Send(
		rs.GetClient().R().
			SetAuthToken(os.Getenv("APPLICATION_ALL_ACCESS_TOKEN")).
			SetQueryParam("email", strings.ToLower(requestBody.Email)),
		"https://" + os.Getenv("APPLICATION_DOMAIN") + "/api/v2/users-by-email",
		resty.GET,
	)

	userInBytes, err := rs.CheckResponse(userExistsCheck, userExistsCheckError, http.StatusOK, resty.AUTH0)
	if err != nil {
		return nil, errors.New("Error in checking of duplicate user.")
	}

	duplicatedUser := []userDto.CreateUserResponseModel{}
	if err := json.Unmarshal(userInBytes, &duplicatedUser); err != nil {
		return nil, errors.New("Error in unmarshalling the users data")
	}

	if len(duplicatedUser) > 0 {
		return &duplicatedUser[0], errors.New("The user alredy exists, skipping the user creation process")
	}

	userCreationReq, err := rs.Send(
		rs.GetClient().R().
			SetAuthToken(os.Getenv("APPLICATION_ALL_ACCESS_TOKEN")).
			SetBody(requestBody),
		"https://" + os.Getenv("APPLICATION_DOMAIN") + "/api/v2/users",
		resty.POST,
	)

	responseInBytes, err := rs.CheckResponse(userCreationReq, err, http.StatusCreated, resty.AUTH0)
	if err != nil {
		return nil, err
	}

	userCreationResponse := userDto.CreateUserResponseModel{}
	if err := json.Unmarshal(responseInBytes, &userCreationResponse); err != nil {
		return nil, err
	}

	return userCreationResponse, nil
}

func ChangePassword(email string) (interface{}, error) {
	rs := resty.GetRestyClient()

	changePassData := userDto.ChangePasswordRequestModel {
		ClientId: "kTWqSEwsBNVgdvh2BJmxzTWACq80dDG6",
		Email: email,
		Connection: "Username-Password-Authentication",
	}

	changePass, err := rs.Send(
		rs.GetClient().R().
			SetAuthToken(os.Getenv("APPLICATION_ALL_ACCESS_TOKEN")).
			SetBody(changePassData),
		"https://" + os.Getenv("APPLICATION_DOMAIN") + "/dbconnections/change_password",
		resty.POST,
	)

	resInBytes, err := rs.CheckResponse(changePass, err, http.StatusOK, resty.AUTH0)
	if err != nil {
		return nil, err
	}

	res := userDto.ChangePasswordResponseModel(string(resInBytes))
	return &res, nil
}