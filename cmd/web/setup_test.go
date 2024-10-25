package main

import (
	"go_breeders/models"

	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		Models: *models.New(nil),
	}

	os.Exit(m.Run())
}
