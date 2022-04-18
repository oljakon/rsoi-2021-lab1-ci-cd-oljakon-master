package main

import (
	"log"
	"net/http"
	"os"
	"rsoi-lab/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")

	r := handlers.Router()

	log.Fatal(http.ListenAndServe(":"+port, r))
}
