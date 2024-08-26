package user_infra

type UserDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
}
