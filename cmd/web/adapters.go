package main

import (
	"encoding/json"
	"encoding/xml"
	"go_breeders/models"
	"io"
	"net/http"
)

type CatBreedsinterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
}

type RemoteService struct {
	Remote CatBreedsinterface
}

func (rs *RemoteService) GetAllCatBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

type JSONBackend struct{}

func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed

	err = json.Unmarshal(body, &breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil

}

type XMLbackEnd struct{}

func (xb *XMLbackEnd) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds

	err = xml.Unmarshal(body, &breeds)

	if err != nil {
		return nil, err
	}

	return breeds.Breeds, nil

}
