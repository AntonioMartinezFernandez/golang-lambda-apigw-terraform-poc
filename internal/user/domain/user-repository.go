package user_domain

type UserRepository interface {
	Find(id string) (*User, error)
	Save(user User) error
}
