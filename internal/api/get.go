package api

import (
	"database/sql"
	"errors"
	"log"

	"rsoi-lab/internal/db"
	"rsoi-lab/models"
)

const getPersonQuery = `SELECT * FROM persons WHERE id=$1`

func (r *PersonRepository) GetPerson(id int) (models.Person, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var person models.Person

	row := r.db.QueryRow(getPersonQuery, id)

	err := row.Scan(&person.ID, &person.Name, &person.Address, &person.Work, &person.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return person, err
		}
	}

	return person, nil
}

const getAllPersonsQuery = `SELECT * FROM persons`

func (r *PersonRepository) GetAllPersons() ([]models.Person, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var persons []models.Person

	rows, err := r.db.Query(getAllPersonsQuery)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		err = rows.Scan(&person.ID, &person.Name, &person.Address, &person.Work, &person.Age)
		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}

		persons = append(persons, person)
	}

	return persons, err
}
