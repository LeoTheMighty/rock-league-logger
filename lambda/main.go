package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"strconv"
	"time"
)

type LoggerRequestEvent struct {
	// create, update, and delete params
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`

	// query params
	StartAt string `json:"start_at"`
	Limit   int64  `json:"limit"`
}

type Log struct {
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func handleResponse(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: body, StatusCode: 200}
}

func handleError(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}
}

func handlePost(event LoggerRequestEvent) (events.APIGatewayProxyResponse, error) {
	svc := initDynamoDB()

	createdAt := time.Now().UTC().Format(time.RFC3339)
	err := createLog(svc, event.UserID, event.Content, createdAt)
	if err != nil {
		return handleError(err), nil
	}

	createdLog := &Log{UserID: event.UserID, Content: event.Content, CreatedAt: createdAt}
	jsonResponse, err := json.Marshal(createdLog)
	if err != nil {
		return handleError(err), nil
	}

	return handleResponse(string(jsonResponse)), nil
}

func handlePatch(event LoggerRequestEvent) (events.APIGatewayProxyResponse, error) {
	svc := initDynamoDB()

	result, err := updateLog(svc, event.UserID, event.CreatedAt, event.Content)
	if err != nil {
		return handleError(err), nil
	}

	// Extracting the "content" from result.Attributes
	updatedContent := result.Attributes["content"]

	// Prepare the content to be marshalled into JSON
	updatedLog := &Log{UserID: event.UserID, Content: *updatedContent.S, CreatedAt: event.CreatedAt}

	jsonResponse, err := json.Marshal(updatedLog)
	if err != nil {
		return handleError(err), nil
	}

	return handleResponse(string(jsonResponse)), nil
}

func handleDelete(event LoggerRequestEvent) (events.APIGatewayProxyResponse, error) {
	svc := initDynamoDB()

	result, err := deleteLog(svc, event.UserID, event.CreatedAt)
	if err != nil {
		return handleError(err), nil
	}

	return handleResponse(result.String()), nil
}

func handleGet(event LoggerRequestEvent) (events.APIGatewayProxyResponse, error) {
	svc := initDynamoDB()

	result, err := queryLogs(svc, event.UserID, event.StartAt, event.Limit)
	if err != nil {
		return handleError(err), nil
	}

	var logs []Log
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &logs)
	if err != nil {
		return handleError(err), nil
	}

	jsonResponse, err := json.Marshal(logs)
	if err != nil {
		return handleError(err), nil
	}

	return handleResponse(string(jsonResponse)), nil
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST", "PATCH", "DELETE":
		var event LoggerRequestEvent
		err := json.Unmarshal([]byte(request.Body), &event)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}
		log.Default().Printf("Processing \"%s\" request with body: `%s`!", request.HTTPMethod, request.Body)

		switch request.HTTPMethod {
		case "POST":
			return handlePost(event)
		case "PATCH":
			return handlePatch(event)
		case "DELETE":
			return handleDelete(event)
		default:
			return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Invalid HTTP method: %s", request.HTTPMethod), StatusCode: 400}, nil
		}
	case "GET":
		log.Default().Printf("Processing GET request with query parameters: %v", request.QueryStringParameters)

		limit, err := strconv.ParseInt(request.QueryStringParameters["limit"], 10, 64)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}

		// We need to convert query params to LoggerRequestEvent struct
		event := LoggerRequestEvent{
			UserID:  request.QueryStringParameters["user_id"],
			StartAt: request.QueryStringParameters["start_at"],
			Limit:   limit,
		}
		return handleGet(event)
	default:
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Invalid HTTP method: %s", request.HTTPMethod), StatusCode: 400}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
