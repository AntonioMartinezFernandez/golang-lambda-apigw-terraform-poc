package user_domain

import "context"

type UserRepository interface {
	Find(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
