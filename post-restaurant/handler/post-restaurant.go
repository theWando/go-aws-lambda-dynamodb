package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/theWando/go-aws-lambda-dynamodb/lambda-lib/logger"
	"github.com/theWando/go-aws-lambda-dynamodb/lambda-lib/model"
	"go.uber.org/zap"
)

type (
	DynamoClient interface {
		TransactWriteItems(ctx context.Context, params *dynamodb.TransactWriteItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error)
	}
	Handler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
)

var (
	failedToSaveItem = events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Sorry! we couldn't save the restaurant"}

	successResponse = events.APIGatewayProxyResponse{StatusCode: http.StatusCreated}
)

func PostRestaurant(db DynamoClient) Handler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var restaurant model.Restaurant
		if err := json.Unmarshal([]byte(request.Body), &restaurant); err != nil {
			logger.GetLoggerFromContext(ctx).Error("failed to json unmarshal restaurant", zap.Error(err))
			return failedToSaveItem, err
		}

		av, err := attributevalue.MarshalMap(restaurant)
		if err != nil {
			logger.GetLoggerFromContext(ctx).Error("failed to marshalMap restaurant", zap.Error(err))
			return failedToSaveItem, err
		}
		items, err := db.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
			TransactItems: []types.TransactWriteItem{
				{
					Put: &types.Put{
						TableName: aws.String("restaurants"),
						Item:      av,
					},
				},
			},
		})
		if err != nil {
			logger.GetLoggerFromContext(ctx).Error("failed to create restaurant", zap.Error(err))
			return failedToSaveItem, err
		}
		logger.GetLoggerFromContext(ctx).Debug("saved", zap.Any("saved-item", items))
		return successResponse, nil
	}
}
