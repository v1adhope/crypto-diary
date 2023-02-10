package dto

type SignRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=32,min=8"`
}

type RefreshToken struct {
	Token string `json:"refreshToken"`
}
