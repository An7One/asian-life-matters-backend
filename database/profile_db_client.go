package database

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	model "github.com/zea7ot/web_api_aeyesafe/model/user"
)

const tableName string = "aeyesafe_user_profile"

// ProfileDBClient returns a client to the dynamo database
type ProfileDBClient struct {
	client *DBClient
}

// NewProfileClient returns a ProfileClient implementation
func NewProfileClient(client *DBClient) *ProfileDBClient {
	return &ProfileDBClient{
		client: client,
	}
}

// GetOneProfileByPhoneNumber returns a Profile by his/her phone number
func (s *ProfileDBClient) GetOneProfileByPhoneNumber(phoneNumber string) (*model.Profile, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"phoneNumber": {
				N: aws.String(phoneNumber),
			},
		},
	}

	res, err := s.client.svc.GetItem(input)
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
func (s *ProfileDBClient) AddOneProfile(p *model.Profile) (*model.Profile, error) {
	return p, nil
}

// UpdateOneProfile updates a Profile with the input Profile
func (s *ProfileDBClient) UpdateOneProfile(p *model.Profile) (*model.Profile, error) {
	return p, nil
}
