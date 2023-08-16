package vo

import (
	"errors"
	"math/rand"
	"regexp"
	"time"
)

var InvalidUserIdError = errors.New("invalid user id")

const (
	_charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	_patternUserId = `^[a-zA-Z0-9][a-zA-Z0-9_]*$`
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
	match, _ := regexp.MatchString(_patternUserId, userId)
	return match
}

func checkUserIdLength(userId string) bool {
	return len(userId) >= 3 && len(userId) <= 20
}

// 有効なuserIdを生成する
func GenerateUserId() UserId {
	return UserId{value: generateRandomString(10)}
}

func generateRandomString(length int) string {
	s := ""
	if length > 0 {
		for i := 0; i < length; i++ {
			s += string(_charset[seededRand.Intn(len(_charset))])
		}
	}
	return s
}

func (u UserId) Value() string {
	return u.value
}

func (u UserId) Equal(o UserId) bool {
	return u.value == o.value
}
