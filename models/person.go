package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"testing"
)

// Person table
type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Work    string `json:"work"`
	Age     int    `json:"age"`
}

func (p *Person) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Address, validation.Required),
		validation.Field(&p.Work, validation.Required),
		validation.Field(&p.Age, validation.Required),
	)
}

func TestPerson(t *testing.T) *Person {
	t.Helper()

	return &Person{
		Name:    "test",
		Address: "test",
		Work:    "test",
		Age:     2,
	}
}
