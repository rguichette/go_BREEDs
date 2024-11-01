package main

import (
	"go_breeders/config"
	"go_breeders/models"

	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	TestBackend := &TestBackend{}
	testAdapter := &RemoteService{
		Remote: TestBackend,
	}

	testApp = application{
		App:        config.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{
			ID: 1, Breed: "Tomcat", Details: "many details",
		},
	}
	return breeds, nil
}
