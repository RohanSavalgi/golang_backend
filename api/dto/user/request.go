package dto

type CreateUserRequestModel struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Connection string `json:"connection"`
	Password   string `json:"password"`
}
