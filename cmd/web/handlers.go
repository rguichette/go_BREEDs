package main

import (
	"fmt"
	pets "go_breeders/Pets"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, world!")
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	fmt.Println("page hit: ", page)

	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))

}
func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))

}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	dog, err := pets.NewPetFromAbstractFactory("dog")

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusAccepted)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}
func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	cat, err := pets.NewPetFromAbstractFactory("cat")

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusAccepted)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusAccepted)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed").
		SetDescription("a mixed breed").
		SetAge(4).
		SetAgeEstimated(true).Build()
	// Setwi
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, p)

	//create a dog using a buildr pattern

}
func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("british Shorthair").
		SetDescription("a mixed breed").
		SetAge(3).
		SetAgeEstimated(true).Build()
	// Setwi
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, p)

	//create a dog using a buildr pattern

}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	catBreeds, err := app.App.CatService.GetAllCatBreeds()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	//Setup toolbox

	// get species from url
	//Get breed from Url

	//create pet from abstract factory

	//write result as json
}
