package models

type User struct {
	ID int `json:"user_id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	ID int `json:"user_id"`
	Name string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID int `json:"user_id"`
	Email string `json:"email"`
}
