package api

import (
	"log"

	"rsoi-lab/internal/db"
)

const deletePersonQuery = `DELETE FROM persons WHERE id=$1`

func (r *PersonRepository) DeletePerson(id int) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	_, err := r.db.Exec(deletePersonQuery, id)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	return
}
