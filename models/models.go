package models

type User struct {
	UserID int `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
