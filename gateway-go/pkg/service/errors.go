package service

import "errors"

var (
	ErrHashingPassword     = errors.New("error hashing password")
	ErrGenerateUniqueName  = errors.New("error when generate unique name")
	ErrDuplicateEmail      = errors.New("email already exists")
	ErrUserAlreadyVerified = errors.New("user already verified")
	ErrRecordNotfound      = errors.New("record not found")
	ErrInvalidToken        = errors.New("invalid token")
)
