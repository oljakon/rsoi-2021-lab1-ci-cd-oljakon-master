package handlers

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/persons", CreatePersonHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/persons/{id}", GetPersonHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/persons", GetAllPersonsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/persons/{id}", PatchPersonHandler).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/api/v1/persons/{id}", DeletePersonHandler).Methods("DELETE", "OPTIONS")

	return router
}
