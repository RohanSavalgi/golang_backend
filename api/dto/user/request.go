package dto

type CreateUserRequestModel struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Connection string `json:"connection"`
	Password   string `json:"password"`
}

type ChangePasswordRequestModel struct {
	ClientId   string `json:"client_id"`
	Email      string `json:"email"`
	Connection string `json:"connection"`
}

type ChangePasswordResponseModel string