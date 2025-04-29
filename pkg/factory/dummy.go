package factory

import "qdrant/pkg/provider"

// DummyClusterConfig returns a minimal fake ClusterConfig
func DummyClusterConfig(id string) provider.ClusterConfig {
	return provider.ClusterConfig{
		Name:       id,
		Region:     "us-east-1",
		NodeCount:  3,
		NodeType:   "t3.medium",
		K8sVersion: "1.32",
	}
}
