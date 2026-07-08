package ecode

import "errors"

var (
	ErrNullPointer     = errors.New("null reference instance")
	ErrEmptyURL        = errors.New("empty URL")
	ErrNoURL           = errors.New("no URL specified")
	ErrInvalidDataPath = errors.New("invalid data path")
)
