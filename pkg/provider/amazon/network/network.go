package network

import (
	"qdrant/pkg/provider"
)

// Info ...
type Info struct {
	VpcID     string
	SubnetIDs []string
}

// Service ...
type Service struct{}

// New ...
func New() *Service {
	return &Service{}
}

// Create ...
func (s *Service) Create(cfg provider.ClusterConfig) (*Info, error) {
	return &Info{
		VpcID:     "vpc-xxxxxxx",
		SubnetIDs: []string{"subnet-xxxx", "subnet-yyyy"},
	}, nil
}
