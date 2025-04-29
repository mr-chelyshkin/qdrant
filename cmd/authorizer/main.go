package main

import (
	"context"
	"os"
	"qdrant/pkg/auth"

	"qdrant/pkg/utils/logger"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

const tokenSeparator = ":::"

var (
	hmacToken = os.Getenv("HMAC_API_TOKEN")
	jwtSecret = os.Getenv("JWT_SECRET")

	log           = logger.InitLogger()
	authenticator *auth.Authenticator
)

func init() {
	if hmacToken == "" || jwtSecret == "" {
		log.Fatal().Msg("AUTH_TOKEN and JWT_SECRET environment variables must be set")
	}
	authenticator = auth.NewAuthenticator(hmacToken, jwtSecret)
}

func generatePolicy(principalID string, effect string, resource string, context map[string]interface{}) (events.APIGatewayCustomAuthorizerResponse, error) {
	if effect != "Allow" && effect != "Deny" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("invalid effect")
	}
	authResponse := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: principalID,
		Context:     context,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		},
	}
	return authResponse, nil
}

func handler(ctx context.Context, req events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authHeader, ok := req.Headers["x-api-auth"]
	if !ok || authHeader == "" {
		log.Error().Msg("x-api-auth header missing")
		return generatePolicy("", "Deny", req.MethodArn, nil)
	}

	parts := strings.Split(authHeader, tokenSeparator)
	switch len(parts) {
	case 1:
		return handleUserAuth(parts[0], req)
	case 2:
		return handleHmacAuth(parts[0], parts[1], req)
	default:
		log.Error().Msg("Invalid x-api-auth header format")
		return generatePolicy("", "Deny", req.MethodArn, nil)
	}
}

func main() {
	lambda.Start(handler)
}
