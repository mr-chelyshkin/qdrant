package clusterservice

import (
	"qdrant/pkg/factory"
	"qdrant/pkg/provider"
	"qdrant/pkg/qdrant"
	helmservice "qdrant/pkg/service/helm_service"
)

// ClusterService ...
type ClusterService struct {
	helmSvc *helmservice.HelmService
}

// NewClusterService ...
func NewClusterService(helmSvc *helmservice.HelmService) *ClusterService {
	return &ClusterService{
		helmSvc: helmSvc,
	}
}

// CreateAndDeployCluster ...
func (c *ClusterService) CreateAndDeployCluster(cloud string, clusterCfg provider.ClusterConfig, infraOpts provider.CreateOptions) error {
	// Create provider.
	prov, err := factory.NewProvider(cloud, clusterCfg)
	if err != nil {
		return err
	}

	// Create K8S infrastructure.
	infra, err := prov.CreateCluster(clusterCfg, infraOpts)
	if err != nil {
		return err
	}

	// Deploy Qdrant via helm.
	values := qdrant.DefaultEnv(clusterCfg.NodeCount).ToMap()

	return c.helmSvc.Install(
		helmservice.InstallRequest{
			ReleaseName: infra.GetClusterName(),
			ChartPath:   "./charts/qdrant",
			Values:      values,
		},
	)
}
