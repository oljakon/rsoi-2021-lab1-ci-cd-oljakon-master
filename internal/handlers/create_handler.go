package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"rsoi-lab/internal/api"
	"rsoi-lab/models"
)

// CreatePersonHandler handlers POST request
func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	pr := api.PersonRepository{}

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := pr.CreatePerson(person)

	w.Header().Set("Location", fmt.Sprintf("/api/v1/persons/%d", id))
	w.WriteHeader(http.StatusCreated)
	return
}
