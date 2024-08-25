package user_application

import "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"

type SaveUserCommand struct {
	id   string
	data map[string]interface{}
}

func NewSaveUserCommand(userId string, userName string, userBirthdate string) *SaveUserCommand {
	id := utils.NewUlid()
	data := map[string]interface{}{
		"id":        userId,
		"name":      userName,
		"birthdate": userBirthdate,
	}

	return &SaveUserCommand{id: id.String(), data: data}
}

func (c *SaveUserCommand) ID() string {
	return c.id
}

func (c *SaveUserCommand) Data() map[string]interface{} {
	return c.data
}
