package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func updateLog(svc *dynamodb.DynamoDB, userID string, createdAt string, newContent string) (*dynamodb.UpdateItemOutput, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newContent": {
				S: aws.String(newContent),
			},
		},
		TableName: aws.String("logs"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userID),
			},
			"created_at": {
				S: aws.String(createdAt),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set content = :newContent"),
	}

	result, err := svc.UpdateItem(input)

	return result, err
}
