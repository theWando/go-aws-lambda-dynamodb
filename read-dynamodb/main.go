package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/theWando/go-aws-lambda-dynamodb/lambda-lib/state"
	"github.com/theWando/go-aws-lambda-dynamodb/read-dynamodb/handlers"
)

func main() {
	state.Warmup()
	lambda.StartWithContext(state.GetContext(), handlers.NewReadRestaurants(state.GetDBClient()))
}
