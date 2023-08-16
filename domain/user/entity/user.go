package entity

import (
	"github.com/w40141/GoTestSample/domain/user/vo"
)

type User struct {
	id   vo.UserId
	name vo.UserName
	age  vo.UserAge
}

func NewUser(id, firstName, lastName string, age int) (User, error) {
	userId, err := vo.NewUserId(id)
	if err != nil {
		return User{}, err
	}
	userName, err := vo.NewUserName(firstName, lastName)
	if err != nil {
		return User{}, err
	}
	userAge, err := vo.NewUserAge(age)
	return User{id: userId, name: userName, age: userAge}, nil
}

func GenerateNewUser(firstName, lastName string, age int) (User, error) {
	userId := vo.GenerateNewUserId()
	userName, err := vo.NewUserName(firstName, lastName)
	if err != nil {
		return User{}, err
	}
	userAge, err := vo.NewUserAge(age)
	return User{id: userId, name: userName, age: userAge}, nil
}

func (u User) Id() vo.UserId {
	return u.id
}

func (u User) IdValue() string {
	return u.id.Value()
}

func (u User) Name() vo.UserName {
	return u.name
}

func (u User) NameValue() string {
	return u.name.FullName()
}

func (u User) Age() vo.UserAge {
	return u.age
}

func (u User) AgeValue() int {
	return u.age.Value()
}
