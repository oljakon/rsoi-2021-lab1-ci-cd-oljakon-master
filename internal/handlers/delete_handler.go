package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"rsoi-lab/internal/api"
)

// DeletePersonHandler handlers DELETE request
func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	pr := api.PersonRepository{}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int:  %v", err)
	}

	pr.DeletePerson(id)

	w.WriteHeader(http.StatusNoContent)
}
