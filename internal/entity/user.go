package entity

type User struct {
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: remove or replace
// type CreateUserDTO struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
