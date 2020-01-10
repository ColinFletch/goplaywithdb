package model

import "errors"

var (
	// ErrNotFound error when post not in DB
	ErrNotFound = errors.New("item not found")
)
