package entity

import "github.com/pkg/errors"

var (
	DuplicateError = errors.New("object already exists")
)
