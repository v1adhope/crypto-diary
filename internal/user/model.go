package user

type User struct {
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
