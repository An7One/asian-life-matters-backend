package database

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	model "github.com/zea7ot/web_api_aeyesafe/model/user"
)

// ProfileOTPDBClient returns a client to the dynamo database - aeye_user_profile_otp
type ProfileOTPDBClient struct {
	db *dynamodb.DynamoDB
}

// NewProfileOTPDBClient returns a client to the
func NewProfileOTPDBClient(db *dynamodb.DynamoDB) *ProfileOTPDBClient {
	return &ProfileOTPDBClient{
		db: db,
	}
}

// AddOneProfileOTP creates a ProfileOTP database item with the input ProfileOTP and inserts it into the database
func (c *ProfileOTPDBClient) AddOneProfileOTP(otp *model.ProfileOTP) (*model.ProfileOTP, error) {
	av, err := dynamodbattribute.MarshalMap(otp)
	if err != nil {
		logrus.Error("Got an error marshalling ProfileOTP item")
		logrus.Error(err.Error())
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(viper.GetString("database.dynamo.table.userProfileOTP.name")),
	}

	_, err = c.db.PutItem(input)
	if err != nil {
		logrus.Error("Got an error calling PutItem:")
		logrus.Error(err.Error())
		return nil, err
	}

	return otp, err
}

// GetOneProfileOTPByPhoneNumber returns a ProfileOTP by the phone number
func (c *ProfileOTPDBClient) GetOneProfileOTPByPhoneNumber(phoneNumber string) (*model.ProfileOTP, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(viper.GetString("database.dynamo.table.userProfileOTP.name")),
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
		return nil, errors.New("Could not find '" + phoneNumber + ";")
	}

	profileOTP := model.ProfileOTP{}
	err = dynamodbattribute.UnmarshalMap(res.Item, profileOTP)
	if err != nil {
		return nil, err
	}

	return &profileOTP, nil
}

// UpdateOneProfileOTP updates a ProfileOTP with the input ProfileOTP
func (c *ProfileOTPDBClient) UpdateOneProfileOTP(otp *model.ProfileOTP) (*model.ProfileOTP, error) {
	return otp, nil
}
