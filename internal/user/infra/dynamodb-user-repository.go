package user_infra

import (
	"context"
	"time"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var _ user_domain.UserRepository = (*DynamoDbUserRepository)(nil) // Check if DynamoDbUserRepository implements the UserRepository interface

const tableName = "users"

type DynamoDbUserRepository struct {
	dynamodbClient *dynamodb.Client
}

func NewDynamoDbUserRepository(dynamodbClient *dynamodb.Client) *DynamoDbUserRepository {
	return &DynamoDbUserRepository{
		dynamodbClient: dynamodbClient,
	}
}

func (dur *DynamoDbUserRepository) Find(ctx context.Context, userId string) (*user_domain.User, error) {
	userModel := UserDynamoDbModel{Id: userId}
	key, kErr := userModel.GetKey()
	if kErr != nil {
		return nil, kErr
	}
	response, err := dur.dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key: key, TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, err
	} else {
		// Not item found
		if len(response.Item) == 0 {
			return nil, nil
		}
		err = attributevalue.UnmarshalMap(response.Item, &userModel)
		if err != nil {
			return nil, err
		}
	}

	return parseUser(userModel), err
}

func (dur *DynamoDbUserRepository) Save(ctx context.Context, user user_domain.User) error {
	u := UserDynamoDbModel{Id: user.Id(), Name: user.Name(), Birthdate: user.Birthdate().Format(time.RFC3339)}

	item, err := attributevalue.MarshalMap(u)
	if err != nil {
		return err
	}

	_, putItemErr := dur.dynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("users"), Item: item,
	})
	return putItemErr
}

func parseUser(userModel UserDynamoDbModel) *user_domain.User {
	birthdate, _ := time.Parse(time.RFC3339, userModel.Birthdate)
	return user_domain.NewUser(
		userModel.Id,
		userModel.Name,
		birthdate,
	)
}
