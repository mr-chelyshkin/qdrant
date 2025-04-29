package main

import (
	"runtime/debug"

	"qdrant/pkg/api"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	debug.SetGCPercent(500)
}

func main() {
	lambda.Start(
		api.NewLambda(
			api.Config{
				EnableRequestLogging: true,
			},
			map[string]api.HandleFunc{
				"GET:/v1/qdrant/access": handleClusterTokenGet,
			},
		).Handle,
	)
}
