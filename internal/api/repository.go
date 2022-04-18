package api

import (
	"database/sql"

	"rsoi-lab/models"
)

type Repository interface {
	CreatePerson(person models.Person) int
	DeletePerson(id int)
	GetPerson(id int) (models.Person, error)
	GetAllPersons() ([]models.Person, error)
	PatchPerson(id int, person models.Person) (models.Person, error)
}

type PersonRepository struct {
	db *sql.DB
}
