package user_infra

type UpdateUserDto struct {
	Id        string  `json:"id"`
	Name      *string `json:"name"`
	Birthdate *string `json:"birthdate"`
}
