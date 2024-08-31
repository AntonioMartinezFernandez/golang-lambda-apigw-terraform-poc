package user_infra

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserDynamoDbModel struct {
	Id        string `dynamodbav:"Id"`
	Name      string `dynamodbav:"Name"`
	Birthdate string `dynamodbav:"Birthdate"`
}

// GetKey returns the composite primary key of the user in a format that can be
// sent to DynamoDB.
func (user UserDynamoDbModel) GetKey() (map[string]types.AttributeValue, error) {
	id, err := attributevalue.Marshal(user.Id)
	if err != nil {
		return map[string]types.AttributeValue{}, err
	}
	return map[string]types.AttributeValue{"Id": id}, nil
}
