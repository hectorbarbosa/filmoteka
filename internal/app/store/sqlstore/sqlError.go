package sqlstore

import (
	"errors"
)

// func getSQLState(err error) string {
// 	type checker interface {
// 		SQLState() string
// 	}
// 	pe := err.(checker)
// 	log.Println("SQLState:", pe.SQLState())
// 	return pe.SQLState()
// }

var (
	ErrResourceNotFound = errors.New("resource not found")
	// ErrResourceNotCreated = errors.New("resource not created")
	ErrUniqueConstraints = errors.New("unique constraints violation")
	ErrValidation        = errors.New("validation error")
)
