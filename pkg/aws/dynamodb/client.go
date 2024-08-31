package dynamo_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewClient(awsRegion string, dynamoDbEndpoint string, debug bool) *dynamodb.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID {
			return aws.Endpoint{
				URL:           dynamoDbEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	cfg := *aws.NewConfig()
	var err error
	if debug {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(awsRegion),
			config.WithEndpointResolverWithOptions(customResolver),
			// Enable debugging:
			config.WithClientLogMode(aws.LogRetries|aws.LogRequest|aws.LogRequestWithBody),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(awsRegion),
			config.WithEndpointResolverWithOptions(customResolver),
		)
	}

	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}
