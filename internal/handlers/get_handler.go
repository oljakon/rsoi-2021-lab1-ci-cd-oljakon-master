package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"rsoi-lab/internal/api"
)

// GetPersonHandler handlers GET one person request
func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	pr := api.PersonRepository{}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int:  %v", err)
	}

	person, err := pr.GetPerson(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(404)
			return
		}
		log.Fatalf("Unable to get person: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetAllPersonsHandler handlers GET all persons request
func GetAllPersonsHandler(w http.ResponseWriter, r *http.Request) {
	pr := api.PersonRepository{}

	persons, err := pr.GetAllPersons()
	if err != nil {
		log.Fatalf("Unable to get all persons: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(persons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
