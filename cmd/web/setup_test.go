package main

import (
	"go_breeders/adapters"
	"go_breeders/config"

	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{
		Remote: testBackend,
	}

	testApp = application{
		App: config.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}

// type TestBackend struct{}

// func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
// 	breeds := []*models.CatBreed{
// 		&models.CatBreed{
// 			ID: 1, Breed: "Tomcat", Details: "many details",
// 		},
// 	}
// 	return breeds, nil
// }
