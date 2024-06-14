package repository

import (
	"errors"

	"github.com/Masterminds/squirrel"
)

var (
	PlaceHolder   = squirrel.Dollar
	ErrBuildQuery = errors.New("error building query")
	ErrExecQuery  = errors.New("error executing")
)
