package user_application

import "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"

type UpdateUserCommand struct {
	id   string
	data map[string]interface{}
}

func NewUpdateUserCommand(userId string, userName *string, userBirthdate *string) *UpdateUserCommand {
	id := utils.NewUlid()
	data := map[string]interface{}{
		"id":        userId,
		"name":      userName,
		"birthdate": userBirthdate,
	}

	return &UpdateUserCommand{id: id.String(), data: data}
}

func (c *UpdateUserCommand) ID() string {
	return c.id
}

func (c *UpdateUserCommand) Data() map[string]interface{} {
	return c.data
}
