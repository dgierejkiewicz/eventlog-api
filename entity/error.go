package entity

import "errors"

var (
		//ErrNotFound not found
		ErrNotFound = errors.New("Not found")

		//ErrInvalidEntity invalid entity
		ErrInvalidEntity = errors.New("Invalid entity")
)
