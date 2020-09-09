package models

type (
	UserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	UserResponse struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	User struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
