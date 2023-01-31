package entity

type User struct {
	ID       string `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
