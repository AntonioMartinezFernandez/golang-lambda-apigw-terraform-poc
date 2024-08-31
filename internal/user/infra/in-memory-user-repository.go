package user_infra

import (
	"context"
	"time"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
)

type InMemoryUserRepository struct {
	users map[string]user_domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	users := map[string]user_domain.User{}

	inMemoryUsersRepository := &InMemoryUserRepository{users: users}

	createSomeExampleUsers(inMemoryUsersRepository)

	return inMemoryUsersRepository
}

func (ur *InMemoryUserRepository) Find(_ context.Context, id string) (*user_domain.User, error) {
	user := ur.users[id]
	return &user, nil
}

func (ur *InMemoryUserRepository) Save(_ context.Context, user user_domain.User) error {
	ur.users[user.Id()] = user
	return nil
}

func createSomeExampleUsers(imur *InMemoryUserRepository) {
	ctx := context.Background()
	exampleUsers := []user_domain.User{
		*user_domain.NewUser("01J64TS9923K5WS395CFE1AP25", "Duke Ellington", time.Now()),
		*user_domain.NewUser("01J64TSVKMNEJN1C0R1J15QZDB", "Charlie Parker", time.Now()),
		*user_domain.NewUser("01J64V13D4AHZ61T4MD7Z53BVZ", "Miles Davis", time.Now()),
		*user_domain.NewUser("01J64V18DXAH27PAGZ7MP01127", "Ella Fitzgerald", time.Now()),
		*user_domain.NewUser("01J64V1CV1T1PPMZ5C24MNDMW9", "Louis Armstrong", time.Now()),
	}

	for _, user := range exampleUsers {
		imur.Save(ctx, user)
	}
}
