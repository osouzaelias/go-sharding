package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go-sharding/internal/application/core/domain"
)

type Adapter struct {
	db *dynamodb.Client
}

func NewAdapter(region string) (*Adapter, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				},
			),
		),
		config.WithCredentialsProvider(
			credentials.StaticCredentialsProvider{
				Value: aws.Credentials{
					AccessKeyID:     "fakeMyKeyId",
					SecretAccessKey: "fakeSecretAccessKey",
					Source:          "Hard-coded credentials; values are irrelevant for local DynamoDB",
				},
			},
		),
	)

	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return &Adapter{db: client}, nil
}

func (a Adapter) Create(ctx context.Context, customer domain.Customer, node string) error {
	item, err := attributevalue.MarshalMap(customer)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(node),
		Item:      item,
	}

	_, err = a.db.PutItem(ctx, input)
	return err
}

func (a Adapter) Read(ctx context.Context, customerID, node string) (domain.Customer, error) {
	key, _ := attributevalue.MarshalMap(map[string]string{"id": customerID})

	input := &dynamodb.GetItemInput{
		TableName: aws.String(node),
		Key:       key,
	}

	customer := domain.Customer{}

	output, err := a.db.GetItem(ctx, input)
	if err != nil {
		return customer, err
	}

	if err = attributevalue.UnmarshalMap(output.Item, &customer); err != nil {
		return customer, err
	}

	return customer, nil
}
