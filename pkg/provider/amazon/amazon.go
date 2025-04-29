package amazon

import (
	"qdrant/pkg/provider"
	"qdrant/pkg/provider/amazon/access"
	"qdrant/pkg/provider/amazon/controlplane"
	"qdrant/pkg/provider/amazon/inputs"
	"qdrant/pkg/provider/amazon/network"
	"qdrant/pkg/provider/amazon/nodegroups"
)

// AWSProvider ...
type AWSProvider struct {
	controlPlane *controlplane.Service
	nodeGroups   *nodegroups.Service
	network      *network.Service
	access       *access.Service
}

// New ...
func New(cfg InputConfig) (*AWSProvider, error) {
	cp, err := controlplane.New(inputs.ControlPlaneInputConfig{
		Region:    cfg.Region,
		RoleArn:   cfg.RoleArn,
		SubnetIds: cfg.SubnetIds,
	})
	if err != nil {
		return nil, err
	}

	networkSvc := network.New()
	nodeGroupsSvc := nodegroups.New()
	accessSvc := access.New()

	return &AWSProvider{
		controlPlane: cp,
		network:      networkSvc,
		nodeGroups:   nodeGroupsSvc,
		access:       accessSvc,
	}, nil
}

// CreateCluster ...
func (a *AWSProvider) CreateCluster(cfg provider.ClusterConfig, opts provider.CreateOptions) (provider.InfrastructureInfo, error) {
	var infra InfrastructureInfo
	var err error

	if opts.CreateNetwork {
		infra.NetworkInfo, err = a.network.Create(cfg)
		if err != nil {
			return nil, err
		}
	}

	if opts.CreateControlPlane {
		infra.ControlPlaneInfo, err = a.controlPlane.Create(cfg)
		if err != nil {
			return nil, err
		}
	}

	if opts.CreateNodeGroups {
		nodeGroupInfo, err := a.nodeGroups.Create(cfg)
		if err != nil {
			return nil, err
		}
		infra.NodeGroupsInfo = []*nodegroups.Info{nodeGroupInfo}
	}

	if opts.CreateAccess {
		infra.AccessInfo, err = a.access.Create(cfg)
		if err != nil {
			return nil, err
		}
	}

	return &infra, nil
}

// DeleteCluster ...
func (a *AWSProvider) DeleteCluster(id string) error {
	return a.controlPlane.DeleteCluster(id)
}

// GetClusterInfo ...
func (a *AWSProvider) GetClusterInfo(id string) (provider.ClusterInfo, error) {
	cpInfo, err := a.controlPlane.GetClusterInfo(id)
	if err != nil {
		return nil, err
	}

	return &ClusterInfo{
		ControlPlaneInfo: *cpInfo,
	}, nil
}
