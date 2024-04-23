package db

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
