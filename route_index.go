package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Person is s type.
type Person struct {
	Name string
}

// index lists all food in the database
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	person := Person{Name: "Victor"}

	json, err := json.Marshal(person)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
}
