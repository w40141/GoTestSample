package vo

import (
	"errors"
)

var InvalidAgeError = errors.New("invalid age")

type UserAge struct {
	value int
}

func NewUserAge(value int) (UserAge, error) {
	if !isValidAge(value) {
		return UserAge{}, InvalidAgeError
	}
	return UserAge{value: value}, nil
}

func isValidAge(value int) bool {
	return value >= 0 && value <= 150
}

func (a UserAge) Value() int {
	return a.value
}

func (a UserAge) Equal(o UserAge) bool {
	return a.value == o.value
}
