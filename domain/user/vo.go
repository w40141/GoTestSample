package user

import (
	"errors"
	"math/rand"
	"regexp"
	"time"
)

var (
	InvalidUserIdError = errors.New("invalid user id")
	InvalidNameError   = errors.New("invalid name")
	InvalidAgeError    = errors.New("invalid age")
)

const (
	charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	patternUserId = `^[a-zA-Z0-9][a-zA-Z0-9_]*$`
	patternName   = `^[a-zA-Z]+$`
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type UserId struct {
	value string
}

func NewUserId(value string) (UserId, error) {
	if !isValidUserID(value) {
		return UserId{}, InvalidUserIdError
	}
	return UserId{value: value}, nil
}

// userIdが有効かどうかをチェックする
func isValidUserID(userId string) bool {
	return checkUserIdCharacters(userId) && checkUserIdLength(userId)
}

func checkUserIdCharacters(userId string) bool {
	match, _ := regexp.MatchString(patternUserId, userId)
	return match
}

func checkUserIdLength(userId string) bool {
	return len(userId) >= 3 && len(userId) <= 20
}

// 有効なuserIdを生成する
func GenerateUserID() UserId {
	return UserId{value: generateRandomString(10)}
}

func generateRandomString(length int) string {
	s := ""
	if length > 0 {
		for i := 0; i < length; i++ {
			s += string(charset[seededRand.Intn(len(charset))])
		}
	}
	return s
}

func (u UserId) Value() string {
	return u.value
}

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
	match, _ := regexp.MatchString(patternName, name)
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

type UserAge struct {
	value int
}

func NewAge(value int) (UserAge, error) {
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
