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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type UserId struct {
	value string
}

func NewUserId(value string) (UserId, error) {
	if !isValidUserID(value) && !isValidLengthUserID(value) {
		return UserId{}, InvalidUserIdError
	}
	return UserId{value: value}, nil
}

// userIdが有効かどうかをチェックする
func isValidUserID(userId string) bool {
	pattern := `^[a-zA-Z0-9][a-zA-Z0-9_]*$`
	match, _ := regexp.MatchString(pattern, userId)
	return match
}

func isValidLengthUserID(userId string) bool {
	return len(userId) >= 3 && len(userId) <= 20
}

// 有効なuserIdを生成する
func GenerateUserID() UserId {
	return UserId{value: generateRandomString(10)}
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (u UserId) Value() string {
	return u.value
}

type UserName struct {
	firstName string
	lastName  string
}

func NewUserName(firstName, lastName string) (UserName, error) {
	if !isInValidName(firstName) || !isInValidName(lastName) {
		return UserName{}, InvalidNameError
	}
	return UserName{firstName: firstName, lastName: lastName}, nil
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

func isInValidName(name string) bool {
	pattern := `^[a-zA-Z]*$`
	match, _ := regexp.MatchString(pattern, name)
	return match
}

type Age struct {
	value int
}

func NewAge(value int) (Age, error) {
	if value < 0 || value > 120 {
		return Age{}, InvalidAgeError
	}
	return Age{value: value}, nil
}

func (a Age) Value() int {
	return a.value
}
