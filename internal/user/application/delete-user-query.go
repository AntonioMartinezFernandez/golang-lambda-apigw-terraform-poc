package user_application

import "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"

type DeleteUserCommand struct {
	id   string
	data map[string]interface{}
}

func NewDeleteUserCommand(userId string) *DeleteUserCommand {
	id := utils.NewUlid()
	data := map[string]interface{}{
		"id": userId,
	}

	return &DeleteUserCommand{id: id.String(), data: data}
}

func (c *DeleteUserCommand) ID() string {
	return c.id
}

func (c *DeleteUserCommand) Data() map[string]interface{} {
	return c.data
}
