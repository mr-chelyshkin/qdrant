package provider

// Provider ...
type Provider interface {
	CreateCluster(cfg ClusterConfig, opts CreateOptions) (InfrastructureInfo, error)
	GetClusterInfo(id string) (ClusterInfo, error)
	DeleteCluster(id string) error
}

// ClusterConfig ...
type ClusterConfig struct {
	Name       string
	Region     string
	NodeCount  int
	NodeType   string
	K8sVersion string
}

// CreateOptions ...
type CreateOptions struct {
	CreateNetwork      bool
	CreateControlPlane bool
	CreateNodeGroups   bool
	CreateAccess       bool
}

// InfrastructureInfo ...
type InfrastructureInfo interface {
	GetClusterID() string
	GetClusterName() string
	GetClusterRegion() string
	GetClusterStatus() string
	GetClusterEndpoint() string
	GetKubeconfig() []byte
}

// ClusterInfo ...
type ClusterInfo interface {
	GetClusterID() string
	GetClusterName() string
	GetClusterRegion() string
	GetClusterStatus() string
	GetClusterEndpoint() string
	GetAccessToken() string
	GetKubeconfig() []byte
}
