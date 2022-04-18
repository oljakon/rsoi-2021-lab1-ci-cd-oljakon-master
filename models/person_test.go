package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *Person
		isValid bool
	}{
		{
			name: "valid",
			u: func() *Person {
				return TestPerson(t)
			},
			isValid: true,
		},
		{
			name: "empty name",
			u: func() *Person {
				u := TestPerson(t)
				u.Name = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty address",
			u: func() *Person {
				u := TestPerson(t)
				u.Address = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty work",
			u: func() *Person {
				u := TestPerson(t)
				u.Work = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty age",
			u: func() *Person {
				u := TestPerson(t)
				u.Age = 0

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
