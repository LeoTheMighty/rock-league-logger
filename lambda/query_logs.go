package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func queryLogs(svc *dynamodb.DynamoDB, userID string, startAt string, limit int64) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("logs"),
		Limit:                  aws.Int64(limit),
		KeyConditionExpression: aws.String("user_id = :uid AND created_at < :start"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":uid": {
				S: aws.String(userID),
			},
			":start": {
				S: aws.String(startAt),
			},
		},
		ScanIndexForward: aws.Bool(false), // To order by createdAt in descending order
	}

	result, err := svc.Query(input)
	return result, err
}
