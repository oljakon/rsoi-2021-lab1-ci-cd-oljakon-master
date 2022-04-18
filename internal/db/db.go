package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

//const dsn = "postgresql://program:test@localhost:5432/persons?sslmode=disable"
const dsn = "postgres://azvlcerlczxzhs:cb870393dac502b7f9be38e45d88184ccb607c8c03a6e0ede441babcdc7095da@ec2-3-209-65-193.compute-1.amazonaws.com:5432/detjlotromvfed"

// CreateConnection to persons db
func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
