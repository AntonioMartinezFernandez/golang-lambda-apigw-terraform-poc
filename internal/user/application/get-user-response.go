package user_application

type GetUserResponse struct {
	Id        string `json:"id" jsonapi:"primary,id"`
	UserId    string `json:"user_id" validate:"required" jsonapi:"attr,id"`
	Name      string `json:"name" validate:"required" jsonapi:"attr,name"`
	BirthDate string `json:"birthdate" validate:"required" jsonapi:"attr,birthdate"`
}

func NewGetUserResponse(id string, userId string, name string, birthdate string) GetUserResponse {
	return GetUserResponse{Id: id, UserId: userId, Name: name, BirthDate: birthdate}
}
