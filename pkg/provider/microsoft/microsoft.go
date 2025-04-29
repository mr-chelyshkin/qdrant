package microsoft

import (
	"qdrant/pkg/provider"
)

// AzureProvider ...
type AzureProvider struct{}

// New ...
func New() *AzureProvider {
	return &AzureProvider{}
}

// CreateCluster ...
func (a *AzureProvider) CreateCluster(cfg provider.ClusterConfig, opts provider.CreateOptions) (provider.InfrastructureInfo, error) {
	return &InfrastructureInfo{
		Name:     cfg.Name,
		Region:   cfg.Region,
		Status:   "CREATING",
		Endpoint: "",
	}, nil
}

// DeleteCluster ...
func (a *AzureProvider) DeleteCluster(id string) error {
	return nil
}

// GetClusterInfo ...
func (a *AzureProvider) GetClusterInfo(id string) (provider.ClusterInfo, error) {
	return &ClusterInfo{
		Name:     "dummy-azure-cluster",
		Region:   "westeurope",
		Status:   "RUNNING",
		Endpoint: "https://dummy.azure.endpoint",
	}, nil
}
