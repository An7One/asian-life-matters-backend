package database

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	model "github.com/zea7ot/web_api_aeyesafe/model/user"
)

const tableName string = "aeyesafe_user_profile"

// ProfileDBClient returns a client to the dynamo database
type ProfileDBClient struct {
	db *dynamodb.DynamoDB
}

// NewProfileClient returns a ProfileClient implementation
func NewProfileClient(db *dynamodb.DynamoDB) *ProfileDBClient {
	return &ProfileDBClient{
		db: db,
	}
}

// GetOneProfileByPhoneNumber returns a Profile by his/her phone number
func (c *ProfileDBClient) GetOneProfileByPhoneNumber(phoneNumber string) (*model.Profile, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"phoneNumber": {
				N: aws.String(phoneNumber),
			},
		},
	}

	res, err := c.db.GetItem(input)
	if err != nil {
		return nil, err
	}

	if res.Item == nil {
		return nil, errors.New("Could not find '" + phoneNumber + "'")
	}

	profile := model.Profile{}
	err = dynamodbattribute.UnmarshalMap(res.Item, profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// AddOneProfile creates a Profile with the input Profile and add it to the database
func (c *ProfileDBClient) AddOneProfile(p *model.Profile) (*model.Profile, error) {
	av, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		log.Fatal("Got an error marshalling Profile item")
		log.Fatal(err.Error())
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = c.db.PutItem(input)
	if err != nil {
		log.Fatal("Got an error calling PutItem:")
		log.Fatal(err.Error())
		return nil, err
	}

	return p, err
}

// UpdateOneProfile updates a Profile with the input Profile
func (c *ProfileDBClient) UpdateOneProfile(p *model.Profile) (*model.Profile, error) {
	return p, nil
}
