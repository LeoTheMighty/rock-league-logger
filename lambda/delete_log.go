package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func deleteLog(svc *dynamodb.DynamoDB, userID string, createdAt string) (*dynamodb.DeleteItemOutput, error) {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("logs"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
			"created_at": {
				S: aws.String(createdAt),
			},
		},
		ReturnValues: aws.String("ALL_OLD"),
	}

	result, err := svc.DeleteItem(input)

	return result, err
}
