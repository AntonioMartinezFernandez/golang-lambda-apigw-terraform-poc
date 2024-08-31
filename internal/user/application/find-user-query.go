package user_application

type FindUserQuery struct {
	userId string
}

func NewFindUserQuery(userId string) *FindUserQuery {
	return &FindUserQuery{userId: userId}
}

func (guq FindUserQuery) Data() map[string]interface{} {
	return map[string]interface{}{
		"user_id": guq.userId,
	}
}
