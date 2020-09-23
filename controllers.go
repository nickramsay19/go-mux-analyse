package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my simple website!\n")
}

func MakePeopleByCountryHandler(p People) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// get url variables
		vars := mux.Vars(r)
		country := vars["country"]

		if peopleByCountry, err := p.GetPeopleByCountry(country); err == nil {
			fmt.Fprintf(w, "%v", peopleByCountry)
		} else {
			fmt.Fprintf(w, "err")
		}
	}
}

func MakePeopleByGenderHandler(p People) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// get url variables
		vars := mux.Vars(r)
		gender := vars["gender"]

		if peopleByGender, err := p.GetPeopleByGender(gender); err == nil {
			fmt.Fprintf(w, "%v", peopleByGender)
		} else {
			fmt.Fprintf(w, "err")
		}
	}
}
