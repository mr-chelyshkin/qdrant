package factory

import (
	"errors"
	"fmt"

	"qdrant/pkg/provider"
	"qdrant/pkg/provider/amazon"
	"qdrant/pkg/provider/google"
	"qdrant/pkg/provider/microsoft"
)

// NewProvider ...
func NewProvider(cloud string, cfg provider.ClusterConfig) (provider.Provider, error) {
	switch cloud {
	case "aws":
		prov, err := amazon.New(amazon.InputConfig{
			Region:    cfg.Region,
			RoleArn:   "arn:aws:iam::123456789012:role/EKSClusterRole",
			SubnetIds: nil,
		})
		if err != nil {
			return nil, errors.Join(ErrFailedToCreateProvider, err)
		}
		return prov, nil

	case "gcp":
		return google.New(), nil

	case "azure":
		return microsoft.New(), nil

	default:
		return nil, errors.Join(ErrUnsupportedCloudProvider, fmt.Errorf("cloud: %s", cloud))
	}
}
