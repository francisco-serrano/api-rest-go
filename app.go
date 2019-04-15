package main

import (
	"./src/github.com/gorilla/mux"
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(people)

	if err != nil {
		panic(err)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	var personRet []Person

	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			personRet = append(personRet, item)
			break
		}
	}

	err := json.NewEncoder(w).Encode(personRet)

	if err != nil {
		panic(err)
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person

	params := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&person)

	person.ID = params["id"]
	people = append(people, person)

	err := json.NewEncoder(w).Encode(people)

	if err != nil {
		panic(err)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(people)

	if err != nil {
		panic(err)
	}
}

func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
