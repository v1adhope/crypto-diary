package entity

import "errors"

var (
	//INFO: User
	ErrUserAlreadyExists = errors.New("user with such email already exists")
	ErrUserNotExists     = errors.New("user with such email not exists")

	//INFO: Position
	ErrNoFoundPosition = errors.New("no found position to delete")
	ErrNothingToChange = errors.New("nothing to change")

	//INFO: Private
	ErrWrongPassword       = errors.New("passwords do not match")
	ErrTokenInTheBlocklisk = errors.New("token in the blocklist")
	ErrTokenInvalid        = errors.New("invalid token")
)
