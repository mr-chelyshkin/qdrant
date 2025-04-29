package amazon

import (
	"qdrant/pkg/provider/amazon/access"
	"qdrant/pkg/provider/amazon/controlplane"
	"qdrant/pkg/provider/amazon/network"
	"qdrant/pkg/provider/amazon/nodegroups"
)

// InfrastructureInfo ...
type InfrastructureInfo struct {
	NetworkInfo      *network.Info
	ControlPlaneInfo *controlplane.Info
	NodeGroupsInfo   []*nodegroups.Info
	AccessInfo       *access.Info
}

// GetClusterID ...
func (i InfrastructureInfo) GetClusterID() string {
	return i.ControlPlaneInfo.ID
}

// GetClusterName ...
func (i InfrastructureInfo) GetClusterName() string {
	return i.ControlPlaneInfo.Name
}

// GetClusterRegion ...
func (i InfrastructureInfo) GetClusterRegion() string {
	return i.ControlPlaneInfo.Region
}

// GetClusterStatus ...
func (i InfrastructureInfo) GetClusterStatus() string {
	return i.ControlPlaneInfo.Status
}

// GetClusterEndpoint ...
func (i InfrastructureInfo) GetClusterEndpoint() string {
	return i.ControlPlaneInfo.Endpoint
}

// GetKubeconfig ...
func (i InfrastructureInfo) GetKubeconfig() []byte {
	return nil
}

// ClusterInfo ...
type ClusterInfo struct {
	ControlPlaneInfo controlplane.Info
}

// GetClusterID ...
func (c ClusterInfo) GetClusterID() string {
	return c.ControlPlaneInfo.ID
}

// GetClusterName ...
func (c ClusterInfo) GetClusterName() string {
	return c.ControlPlaneInfo.Name
}

// GetClusterRegion ...
func (c ClusterInfo) GetClusterRegion() string {
	return c.ControlPlaneInfo.Region
}

// GetClusterStatus ...
func (c ClusterInfo) GetClusterStatus() string {
	return c.ControlPlaneInfo.Status
}

// GetClusterEndpoint ...
func (c ClusterInfo) GetClusterEndpoint() string {
	return c.ControlPlaneInfo.Endpoint
}

// GetKubeconfig ...
func (c ClusterInfo) GetKubeconfig() []byte {
	return nil
}

// GetAccessToken ...
func (c ClusterInfo) GetAccessToken() string {
	return "dummy-token"
}
