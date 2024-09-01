package user_infra

type CreateUserDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
}
