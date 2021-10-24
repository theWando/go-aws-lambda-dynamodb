package state

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/theWando/go-aws-lambda-dynamodb/lambda-lib/logger"
	"go.uber.org/zap"
)

var (
	warmedUp bool

	log *zap.Logger

	awsConfig aws.Config

	dynamoDBClient *dynamodb.Client

	appCtx context.Context
)

func Warmup() {
	if warmedUp {
		log.Info("warmed up al ready")
		return
	}
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		panic("failed to get logger")
	}
	log.Info("warming up")
	log.Sugar()
	appCtx = logger.SetLoggerInContext(context.Background(), log)
	awsConfig, err = config.LoadDefaultConfig(appCtx)
	if err != nil {
		log.Panic("failed to load config", zap.Error(err))
	}

	dynamoDBClient = dynamodb.NewFromConfig(awsConfig)
}

func GetDBClient() *dynamodb.Client {
	return dynamoDBClient
}

func GetContext() context.Context {
	return appCtx
}
