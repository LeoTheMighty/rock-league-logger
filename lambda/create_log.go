package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createLog(svc *dynamodb.DynamoDB, userID string, content string, createdAt string) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
			"content": {
				S: aws.String(content),
			},
			"created_at": {
				S: aws.String(createdAt),
			},
		},
		TableName: aws.String("logs"),
	}

	_, err := svc.PutItem(input)
	return err
}
