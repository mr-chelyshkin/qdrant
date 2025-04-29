package openapi

import "qdrant/openapi-interface/gen/qdrantapi"

var (
	DataResponseMessage = func(message string) qdrantapi.ResponseMessage {
		return qdrantapi.ResponseMessage{Message: message}
	}
)
