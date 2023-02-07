package entity

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with such email already exists")
	ErrUserNotExists     = errors.New("user with such email not exists")
	//TODO: looks bad
	ErrWrongPassword = errors.New("passwords do not match")
)
