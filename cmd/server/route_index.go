package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devacto/grobot/data"
)

// index lists all food in the database
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(data.GetAllFoods())

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
}
