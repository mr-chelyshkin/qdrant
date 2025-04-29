package main

import (
	"runtime/debug"

	"qdrant/pkg/api"
	"qdrant/pkg/validator"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	validate *validator.Validator
)

func init() {
	debug.SetGCPercent(500)
	validate = validator.New()
}

func main() {
	lambda.Start(
		api.NewLambda(
			api.Config{
				EnableRequestLogging: true,
			},
			map[string]api.HandleFunc{
				// create
				"POST:/v1/qdrant/cluster": handleClusterPost,
			},
		).Handle,
	)
}
