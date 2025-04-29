package main

import (
	"strconv"

	"qdrant/pkg/auth"

	"github.com/aws/aws-lambda-go/events"
)

func handleHmacAuth(timestamp string, signature string, req events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	if err := authenticator.ValidateHmacRequest(timestamp, signature); err != nil {
		log.Error().Err(err).Msg("Device authentication failed")
		return generatePolicy("", "Deny", req.MethodArn, nil)
	}
	context := map[string]any{
		"permissions": strconv.Itoa(auth.GetPermissionLevel(auth.Admin)),
		"role":        strconv.Itoa(int(auth.Admin)),
		"kind":        strconv.Itoa(int(auth.HMAC)),
	}
	return generatePolicy("device", "Allow", req.MethodArn, context)
}
