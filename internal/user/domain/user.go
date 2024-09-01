package user_domain

import (
	"time"
)

type User struct {
	id        string
	name      string
	birthdate time.Time
}

func NewUser(id string, name string, birthdate time.Time) *User {
	return &User{
		id:        id,
		name:      name,
		birthdate: birthdate,
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Birthdate() time.Time {
	return u.birthdate
}

func (u *User) Update(name *string, birthday *time.Time) {
	if name != nil {
		u.name = *name
	}
	if birthday != nil {
		u.birthdate = *birthday
	}
}
