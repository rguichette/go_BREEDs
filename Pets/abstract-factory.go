package pets

import (
	"errors"
	"fmt"
	"go_breeders/config"
	"go_breeders/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (df *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", df.Pet.Breed.Breed)
}

func (df *dogAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := config.GetInstance()

	breed, _ := app.Models.DogBreed.GetBreedByName(b)
	return &DogFromFactory{Pet: &models.Dog{
		Breed: *breed,
	}}
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

type dogAbstractFactory struct{}

func (df *dogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type catAbstractFactory struct{}

func (cf *catAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (cf *catAbstractFactory) newPetWithBreed(n string) AnimalInterface {
	//get breed for cat

	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory dogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory catAbstractFactory
		cat := catFactory.newPet()

		return cat, nil

	default:
		return nil, errors.New("invalid species supplied")
	}

}

func NewPetWithBreedFromAbstractFactory(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		//return dog with breed imbedded
		var dogFactory dogAbstractFactory

		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		//return Cat with breed embedded

		return &CatFromFactory{}, nil
	default:
		return nil, errors.New("invalid species")
	}
}
