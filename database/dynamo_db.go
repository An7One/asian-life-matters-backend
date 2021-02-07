package database

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DBClient implements database operations for database management
type DBClient struct {
	svc *dynamodb.DynamoDB
}

// GetOneItem returns the item designated
// func (client *DBClient) GetOneItem(item interface{}, tableName string)(res interface{}, error) {
// 	av, err := dynamodbattribute.MarshalMap(item)
// 	if err != nil {
// 		log.Fatal("Error with marshalling the item")
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	input := &dynamodb.GetItemInput{
// 		Key:       av,
// 		TableName: aws.String(tableName),
// 	}

// 	srv := getDynamoDB()
// 	res, err := srv.GetItem(input)
// 	if err != nil {
// 		log.Fatal("Error with getting an item")
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	if res.Item == nil{
// 		log.Fatal("Cannot find the item")
// 		return
// 	}

// 	log.Printf("Sucessfully got an item at %s\n")

// 	err := dynamodbattribute.UnmarshalMap(res.Item, &item)
// 	if err != nil{
// 		log.Fatal("failed to unmarshal Item, %v", err)
// 	}

// 	return item
// }

// CreateOneItem creates one generic item in the dynamo database
// func (client *DBClient) CreateOneItem(item interface{}, tableName string) {
// 	av, err := dynamodbattribute.MarshalMap(item)
// 	if err != nil {
// 		log.Fatal("Error with marshalling new item")
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	input := &dynamodb.PutItemInput{
// 		Item:      av,
// 		TableName: aws.String(tableName),
// 	}

// 	svc := getDynamoDB()
// 	_, err = svc.PutItem(input)
// 	if err != nil {
// 		log.Fatal("Error with creating an item")
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	log.Printf("Successfully added an item at %s\n", time.Now())
// }

// DBConn creates a session connecting to the dyanmo database
func DBConn() *DBClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	return &DBClient{
		svc: svc,
	}
}
