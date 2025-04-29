package main

import (
	"context"
	"encoding/json"
	"net/http"

	"qdrant/openapi-interface"
	"qdrant/openapi-interface/gen/qdrantapi"
	"qdrant/pkg/api"
	"qdrant/pkg/factory"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func handleClusterTokenGet(ctx context.Context, logger zerolog.Logger, _ json.RawMessage, baseParams openapi.QueryParams) (any, *api.HandleError) {
	clusterID := baseParams.GetStringPtr("cluster_id")
	if clusterID == nil || *clusterID == "" {
		return nil, &api.HandleError{
			Status: http.StatusBadRequest,
			Err:    errors.New("missing required query parameter: cluster_id"),
		}
	}

	logger.Info().
		Str("cluster_id", *clusterID).
		Msg("received get access token request")

	// mock aws.
	prov, err := factory.NewProvider("aws", factory.DummyClusterConfig(*clusterID))
	if err != nil {
		return nil, &api.HandleError{
			Status: http.StatusInternalServerError,
			Err:    errors.Wrap(err, "failed to create provider for cluster"),
		}
	}

	clusterInfo, err := prov.GetClusterInfo(*clusterID)
	if err != nil {
		return nil, &api.HandleError{
			Status: http.StatusInternalServerError,
			Err:    errors.Wrap(err, "failed to fetch cluster info"),
		}
	}

	token := clusterInfo.GetAccessToken()

	logger.Info().
		Str("cluster_id", *clusterID).
		Str("token", token).
		Msg("generated access token successfully")

	return qdrantapi.ResponseAccessToken{
		Token: token,
	}, nil
}
