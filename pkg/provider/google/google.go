package google

import (
	"qdrant/pkg/provider"
)

// GCPProvider ...
type GCPProvider struct{}

// New ...
func New() *GCPProvider {
	return &GCPProvider{}
}

// CreateCluster ...
func (g *GCPProvider) CreateCluster(cfg provider.ClusterConfig, opts provider.CreateOptions) (provider.InfrastructureInfo, error) {
	return &InfrastructureInfo{
		Name:     cfg.Name,
		Region:   cfg.Region,
		Status:   "CREATING",
		Endpoint: "",
	}, nil
}

// DeleteCluster ...
func (g *GCPProvider) DeleteCluster(id string) error {
	return nil
}

// GetClusterInfo ...
func (g *GCPProvider) GetClusterInfo(id string) (provider.ClusterInfo, error) {
	return &ClusterInfo{
		Name:     "dummy-cluster",
		Region:   "us-central1",
		Status:   "RUNNING",
		Endpoint: "https://dummy.endpoint",
	}, nil
}
