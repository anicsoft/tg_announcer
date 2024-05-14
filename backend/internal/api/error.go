package api

import "errors"

var (
	ErrDecodeBody = errors.New("failed to decode body")
	ErrParseLoc   = errors.New("failed to parse latitude and longitude")
)
