package user_application

type GetUserQuery struct {
	userId string
}

func NewGetUserQuery(userId string) *GetUserQuery {
	return &GetUserQuery{userId: userId}
}

func (guq GetUserQuery) Data() map[string]interface{} {
	return map[string]interface{}{
		"user_id": guq.userId,
	}
}
