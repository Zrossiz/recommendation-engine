package apperrors

import "errors"

var (
	ErrDBQuery         = errors.New("database query error")
	ErrNotFound        = errors.New("record not found")
	ErrHashPassword    = errors.New("hashing password")
	ErrInvalidPassword = errors.New("password invalid")
	ErrAlreadyExist    = errors.New("already exist")
	ErrUserNotFound    = errors.New("user not found")
)
