package database

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// CreateOneItem creates one generic item in the dynamo database
func CreateOneItem(item interface{}, tableName string) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatal("Error with marshalling new item")
		log.Fatal(err.Error())
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	svc := getDynamoDB()
	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatal("Error with creating an item")
		log.Fatal(err.Error())
		return
	}

	log.Printf("Successfully added an item at %s\n", time.Now())
}

func getDynamoDB() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	return svc
}
