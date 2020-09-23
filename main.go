package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// setup model
	people, err := CreatePeople("data.csv")
	if err != nil {
		log.Fatal("Couldn't read the provided csv.")
	}

	// setup the mux router
	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/country/{country}", MakePeopleByCountryHandler(people))
	r.HandleFunc("/gender/{gender}", MakePeopleByGenderHandler(people))
	http.Handle("/", r)
	
	// serve and check for error
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}