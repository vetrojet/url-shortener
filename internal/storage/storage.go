package storage

import "errors"

var (
	errURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("urn exist")
)
