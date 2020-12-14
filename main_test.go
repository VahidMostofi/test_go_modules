package main

import (
	"testing"

	"github.com/vahidmostofi/test_go_modules/pkg/database"
)

func Test(t *testing.T) {
	db := database.GetUniqueDatabase()
	db.Add("vahid")
	db.Add("vahid")
	db.Add("vahid")
	db.Add("saeed")
	db.Count()
}
