package dto

type CreateUserResponseModel struct {
	UserId  string `json:"user_id"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
}