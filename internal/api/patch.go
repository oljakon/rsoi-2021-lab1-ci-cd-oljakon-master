package api

import (
	"log"

	"rsoi-lab/internal/db"
	"rsoi-lab/models"
)

const patchPersonQuery = `UPDATE persons SET name=$2, address=$3, work=$4, age=$5 WHERE id=$1`

func (r *PersonRepository) PatchPerson(id int, person models.Person) (models.Person, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	oldPerson, err := r.GetPerson(id)
	r.db = db.CreateConnection()
	defer r.db.Close()

	if person.Name == "" && person.Address == "" && person.Work == "" && person.Age == 0 {
		return oldPerson, nil
	}

	if person.Name == "" {
		person.Name = oldPerson.Name
	}

	if person.Address == "" {
		person.Address = oldPerson.Address
	}

	if person.Work == "" {
		person.Work = oldPerson.Work
	}

	if person.Age == 0 {
		person.Age = oldPerson.Age
	}

	_, err = r.db.Exec(patchPersonQuery, id, person.Name, person.Address, person.Work, person.Age)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	return r.GetPerson(id)
}
