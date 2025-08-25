package rsa

import (
	"errors"
)

var (
	ErrRSAInvalidArg error = errors.New("rsa coder error: coder base should be prime and positive")
	ErrRSAIntenal error = errors.New("rsa coder error: internal")
)
