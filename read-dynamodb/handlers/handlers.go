package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.uber.org/zap"
	"read-dynamodb/logger"
)

var (
	// errNon200Response non 200 status code in response
	errNon200Response = errors.New("non 200 Response found")

	failureResponse = events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Oops!"}

	notFountResponse = events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "No restaurants found"}
)

type (
	response struct {
		Restaurants []map[string]types.AttributeValue `json:"restaurants"`
	}

	HandlerFunction func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

	DynamoClient interface {
		Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	}
)

func NewReadRestaurants(db DynamoClient) HandlerFunction {

	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		res, err := db.Scan(ctx, &dynamodb.ScanInput{
			TableName: aws.String("restaurants"),
		})
		if err != nil {
			logger.GetLoggerFromContext(ctx).Error("couldn't get res from restaurants", zap.Error(err))
			return failureResponse, err
		}

		if len(res.Items) == 0 {
			return notFountResponse, nil
		}

		bytes, err := json.Marshal(response{res.Items})
		if err != nil {
			logger.GetLoggerFromContext(ctx).Error("failed to marshall items from dynamodb")
		}
		return events.APIGatewayProxyResponse{
			Body:       string(bytes),
			StatusCode: 200,
		}, nil
	}
}
