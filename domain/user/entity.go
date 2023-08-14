package user

type User struct {
	id   UserId
	name UserName
	age  UserAge
}

func NewUser(id, firstName, lastName string, age int) (User, error) {
	userId, err := NewUserId(id)
	if err != nil {
		return User{}, err
	}
	userName, err := NewUserName(firstName, lastName)
	if err != nil {
		return User{}, err
	}
	userAge, err := NewAge(age)
	return User{id: userId, name: userName, age: userAge}, nil
}

func (u User) Id() UserId {
	return u.id
}

func (u User) IdValue() string {
	return u.id.Value()
}

func (u User) Name() UserName {
	return u.name
}

func (u User) NameValue() string {
	return u.name.FullName()
}

func (u User) Age() UserAge {
	return u.age
}

func (u User) AgeValue() int {
	return u.age.Value()
}
