package api

import (
	"log"

	"rsoi-lab/internal/db"
	"rsoi-lab/models"
)

const createQuery = `INSERT INTO persons (name, address, work, age) VALUES ($1, $2, $3, $4) RETURNING id`

func (r *PersonRepository) CreatePerson(person models.Person) int {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var id int
	err := r.db.QueryRow(createQuery, person.Name, person.Address, person.Work, person.Age).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	return id
}
