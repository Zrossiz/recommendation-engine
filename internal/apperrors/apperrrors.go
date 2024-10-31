package apperrors

import "errors"

var (
	ErrDBQuery      = errors.New("database query error")
	ErrNotFound     = errors.New("record not found")
	ErrHashPassword = errors.New("error hashing password")
)
