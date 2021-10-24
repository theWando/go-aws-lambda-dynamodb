package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"read-dynamodb/handlers"
	"read-dynamodb/state"
)

func main() {
	state.Warmup()
	lambda.StartWithContext(state.GetContext(), handlers.NewReadRestaurants(state.GetDBClient()))
}
