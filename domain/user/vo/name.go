package vo

import (
	"errors"
	"regexp"
)

var InvalidNameError = errors.New("invalid name")

const (
	_patternName = `^[a-zA-Z]+$`
)

type UserName struct {
	firstName string
	lastName  string
}

func NewUserName(firstName, lastName string) (UserName, error) {
	if !isValidName(firstName) || !isValidName(lastName) {
		return UserName{}, InvalidNameError
	}
	return UserName{firstName: firstName, lastName: lastName}, nil
}

func isValidName(name string) bool {
	match, _ := regexp.MatchString(_patternName, name)
	return match
}

func (u UserName) FirstName() string {
	return u.firstName
}

func (u UserName) LastName() string {
	return u.lastName
}

func (u UserName) FullName() string {
	return u.firstName + " " + u.lastName
}

func (u UserName) Equal(o UserName) bool {
	return u.firstName == o.firstName && u.lastName == o.lastName
}
