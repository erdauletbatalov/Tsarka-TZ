package domain

import "errors"

var (
	ErrNotValid      = errors.New("not valid data")
	ErrWrongPassword = errors.New("wrong password")
	ErrExists        = errors.New("exists")
)
