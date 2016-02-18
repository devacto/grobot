package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/devacto/grobot/data"
)

// search lists all food in the database
func search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	terms := ""
	if (len(r.Form["q"]) == 1) {
		terms = r.Form["q"][0]
	}

	json, err := json.Marshal(data.SearchFoods(terms))

	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
}
