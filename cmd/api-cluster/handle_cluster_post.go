package main

import (
	"context"
	"encoding/json"
	"net/http"

	"qdrant/openapi-interface"
	"qdrant/openapi-interface/gen/qdrantapi"
	"qdrant/pkg/api"
	"qdrant/pkg/factory"
	"qdrant/pkg/helm/client"
	"qdrant/pkg/provider"
	"qdrant/pkg/qdrant"
	helmservice "qdrant/pkg/service/helm_service"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func handleClusterPost(ctx context.Context, logger zerolog.Logger, raw json.RawMessage, baseParams openapi.QueryParams) (any, *api.HandleError) {
	var req qdrantapi.RequestCreateCluster

	if err := json.Unmarshal(raw, &req); err != nil {
		return nil, &api.HandleError{
			Status: http.StatusBadRequest,
			Err:    errors.Wrap(err, "failed to parse RequestCreateCluster"),
		}
	}
	if err := validate.ValidateStruct(&req); err != nil {
		return nil, &api.HandleError{Status: http.StatusBadRequest, Err: err}
	}

	logger.Info().
		Str("name", req.Name).
		Int("node_count", int(req.NodeCount)).
		Str("node_type", req.NodeType).
		Str("k8s_version", req.K8sVersion).
		Msg("received create cluster request")

	// Create Provider.
	prov, err := factory.NewProvider(req.Provider, provider.ClusterConfig{
		Name:       req.Name,
		Region:     "us-east-1",
		NodeCount:  int(req.NodeCount),
		NodeType:   req.NodeType,
		K8sVersion: req.K8sVersion,
	})
	if err != nil {
		return nil, &api.HandleError{
			Status: http.StatusInternalServerError,
			Err:    errors.Wrap(err, "failed to create cloud provider"),
		}
	}

	// Create cluster.
	infra, err := prov.CreateCluster(provider.ClusterConfig{
		Name:       req.Name,
		Region:     "us-east-1",
		NodeCount:  int(req.NodeCount),
		NodeType:   req.NodeType,
		K8sVersion: req.K8sVersion,
	}, provider.CreateOptions{
		CreateNetwork:      true,
		CreateControlPlane: true,
		CreateNodeGroups:   true,
		CreateAccess:       true,
	})
	if err != nil {
		return nil, &api.HandleError{
			Status: http.StatusInternalServerError,
			Err:    errors.Wrap(err, "failed to create cluster infrastructure"),
		}
	}

	// Helm install Qdrant.
	helmClient, err := client.NewClient()
	if err != nil {
		return nil, &api.HandleError{
			Status: http.StatusInternalServerError,
			Err:    errors.Wrap(err, "failed to create cluster infrastructure"),
		}
	}
	helmSvc := helmservice.New(helmClient)

	err = helmSvc.Install(helmservice.InstallRequest{
		ReleaseName: "qdrant",
		ChartPath:   "./charts/qdrant",
		Values:      qdrant.DefaultEnv(int(req.NodeCount)).ToMap(),
	})
	if err != nil {
		return nil, &api.HandleError{
			Status: http.StatusInternalServerError,
			Err:    errors.Wrap(err, "failed to create cluster infrastructure"),
		}
	}

	logger.Info().Msg("Qdrant installed successfully")
	return qdrantapi.ResponseClusterInfo{
		Id:       infra.GetClusterID(),
		Name:     infra.GetClusterName(),
		Status:   infra.GetClusterStatus(),
		Endpoint: infra.GetClusterEndpoint(),
	}, nil
}
