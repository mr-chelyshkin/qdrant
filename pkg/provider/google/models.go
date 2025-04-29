package google

// InfrastructureInfo ...
type InfrastructureInfo struct {
	Name     string
	Region   string
	Status   string
	Endpoint string
}

// GetClusterID ...
func (i *InfrastructureInfo) GetClusterID() string {
	return i.Name
}

// GetClusterName ...
func (i *InfrastructureInfo) GetClusterName() string {
	return i.Name
}

// GetClusterRegion ...
func (i *InfrastructureInfo) GetClusterRegion() string {
	return i.Region
}

// GetClusterStatus ...
func (i *InfrastructureInfo) GetClusterStatus() string {
	return i.Status
}

// GetClusterEndpoint ...
func (i *InfrastructureInfo) GetClusterEndpoint() string {
	return i.Endpoint
}

// GetKubeconfig ...
func (i *InfrastructureInfo) GetKubeconfig() []byte {
	return nil
}

// ClusterInfo ...
type ClusterInfo struct {
	Name     string
	Region   string
	Status   string
	Endpoint string
}

// GetClusterID ...
func (c *ClusterInfo) GetClusterID() string {
	return c.Name // Пока без явного ID
}

// GetClusterName ...
func (c *ClusterInfo) GetClusterName() string {
	return c.Name
}

// GetClusterRegion ...
func (c *ClusterInfo) GetClusterRegion() string {
	return c.Region
}

// GetClusterStatus ...
func (c *ClusterInfo) GetClusterStatus() string {
	return c.Status
}

// GetClusterEndpoint ...
func (c *ClusterInfo) GetClusterEndpoint() string {
	return c.Endpoint
}

// GetKubeconfig ...
func (c *ClusterInfo) GetKubeconfig() []byte {
	return nil
}

// GetAccessToken ...
func (c ClusterInfo) GetAccessToken() string {
	// In production, it would be generated dynamically using AWS STS or via aws-iam-authenticator.
	return "dummy-token"
}
