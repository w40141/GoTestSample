package user

type User struct {
	id   UserId
	name UserName
	age  Age
}

func NewUser(id UserId, name UserName, age Age) User {
	return User{id: id, name: name, age: age}
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

func (u User) Age() Age {
	return u.age
}

func (u User) AgeValue() int {
  return u.age.Value()
}
