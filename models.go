package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"errors"
	"strconv"
)

type Person struct {
	id 		int 	`csv:id`
	name 	string 	`csv:name`
	country string 	`csv:country`
	gender 	string 	`csv:gender`
}

type People []Person

func CreatePeople(fileName string) (People, error) {

	// declare new people object
	var newPeople People

	// get data from csv
	if lines, err := ReadCsv(fileName); err == nil {
		for _, line := range lines {
			// check if id can be read
			if id, err := strconv.Atoi(line[0]); err == nil {
				newPeople = append(newPeople, Person{
					id: id,
					name: line[1],
					country: line[2],
					gender: line[3],
				})
			}
		}
	} else {
		return newPeople, err
	}
	return newPeople, nil
}

func (p People) GetPersonN(n int) (Person, error) {
	if len(p) > n {
		return p[n], nil
	} else {
		return Person{-1, "", "", ""}, errors.New(fmt.Sprintf("There is no person with id %d", n))
	}
}

func (p People) GetPeopleByCountry(country string) (People, error) {
	var peopleByCountry People
	for _, person := range p {
		if person.country == country {
			peopleByCountry = append(peopleByCountry, person)
		}
	}

	return peopleByCountry, nil
}

func (p People) GetPeopleByGender(gender string) (People, error) {
	var peopleByGender People
	for _, person := range p {
		if person.gender == gender {
			peopleByGender = append(peopleByGender, person)
		}
	}

	return peopleByGender, nil
}

func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

	f.Close()
    return lines, nil
}
