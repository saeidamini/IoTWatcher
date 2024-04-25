package db

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

type DynamoDBInstance struct {
	Client *dynamodb.DynamoDB
	table  string
}

func NewDynamoDBInstance(client *dynamodb.DynamoDB, table string) (*DynamoDBInstance, error) {
	if client == nil {
		return nil, errors.New("dynamoDB client is required")
	}

	if table == "" {
		return nil, errors.New("table name is required")
	}

	return &DynamoDBInstance{
		Client: client,
		table:  table,
	}, nil
}

func (db *DynamoDBInstance) GetTableName() string {
	return db.table
}

func CreateDynamoDBInstance() *DynamoDBInstance {
	config := &aws.Config{
		Region:   aws.String(os.Getenv("REGION")),
		Endpoint: aws.String(os.Getenv("ENDPOINT")),
	}
	sess := session.Must(session.NewSession(config))

	dynamoDBClient := dynamodb.New(sess)
	dbInstance, err := NewDynamoDBInstance(dynamoDBClient, os.Getenv("DYNAMODB_TABLE"))
	if err != nil {
		log.Fatalf("failed to create DynamoDB instance: %v", err)
	}
	return dbInstance
}
