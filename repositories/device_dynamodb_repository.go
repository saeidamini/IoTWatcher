package repositories

import (
	"cmp"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"simple-api-go/db"
	"simple-api-go/models"
	"simple-api-go/utils"
)

type DeviceDynamoRepository struct {
	db *db.DynamoDBInstance
}

func NewDynamoDeviceService(db *db.DynamoDBInstance) *DeviceDynamoRepository {
	return &DeviceDynamoRepository{
		db: db,
	}
}

func (d *DeviceDynamoRepository) CreateDevice(device *models.Device) (*models.Device, error) {
	av, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.db.GetTableName()),
	}

	_, err = d.db.Client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (d *DeviceDynamoRepository) GetDevice(id string) (*models.Device, error) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(d.db.GetTableName()),
	}

	result, err := d.db.Client.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, utils.ErrDeviceNotFound
	}

	device := &models.Device{}
	err = dynamodbattribute.UnmarshalMap(result.Item, device)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (d *DeviceDynamoRepository) UpdateDevice(id string, updatedDevice *models.Device) (*models.Device, error) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return nil, err
	}

	existingDevice, err := d.GetDevice(id)
	if err != nil {
		return nil, errors.New("device not found")
	}

	expressionAttributeNames := map[string]*string{
		"#N":  aws.String("name"),
		"#DM": aws.String("deviceModel"),
		"#NT": aws.String("note"),
		"#S":  aws.String("serial"),
	}

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":name":        &dynamodb.AttributeValue{S: aws.String(cmp.Or(updatedDevice.Name, existingDevice.Name))},
		":deviceModel": &dynamodb.AttributeValue{S: aws.String(updatedDevice.DeviceModel)},
		":note":        &dynamodb.AttributeValue{S: aws.String(updatedDevice.Note)},
		":serial":      &dynamodb.AttributeValue{S: aws.String(updatedDevice.Serial)},
	}

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(d.db.GetTableName()),
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
		UpdateExpression:          aws.String("SET #N = :name, #DM = :deviceModel, #NT = :note, #S = :serial"),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err = d.db.Client.UpdateItem(input)
	if err != nil {
		return nil, err
	}

	return updatedDevice, nil
}

func (d *DeviceDynamoRepository) DeleteDevice(id string) error {
	key, err := dynamodbattribute.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(d.db.GetTableName()),
	}

	_, err = d.db.Client.DeleteItem(input)
	if err != nil {
		if err.Error() == "item not found" {
			return utils.ErrDeviceNotFound
		}
		return err
	}

	return nil
}
