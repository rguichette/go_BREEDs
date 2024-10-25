package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TextApplication_GetAllDogBreedsJSON(t *testing.T) {

	//create a request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	//create res recorder
	rr := httptest.NewRecorder()
	// create Handler
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)
	//serve handler
	handler.ServeHTTP(rr, req)
	//check res

	if rr.Code != http.StatusOK {
		t.Errorf("wrong res code; got %d, wanted 200", rr.Code)
	}

}
