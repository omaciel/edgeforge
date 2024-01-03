package clients

import "errors"

var (
	ErrNoUsernameProvided = errors.New("no username provided")
	ErrNoPasswordProvided = errors.New("no password provided")
	ErrNoBaseUrlProvided  = errors.New("no baseURL provided")
)
