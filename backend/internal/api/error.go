package api

import "errors"

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrDecodeBody   = errors.New("failed to decode body")
	ErrUserNotFound = errors.New("user not found")
	ErrNotAllowed   = errors.New("not allowed")
)
