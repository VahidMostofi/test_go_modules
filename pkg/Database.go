package database

import (
	"github.com/vahidmostofi/test_go_modules/internal/unique"
)

// Database store items in different ways for fun
type Database interface {
	Add(interface{}) error
	Count() int
}

// GetUniqueDatabase returns a unique database
func GetUniqueDatabase() Database {
	ud := unique.UniqueDatabase{}

	return ud
}
