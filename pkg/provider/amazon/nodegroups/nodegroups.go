package nodegroups

import (
	"qdrant/pkg/provider"
)

// Info ...
type Info struct {
	Name string
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
		Name: "nodegroup-default",
	}, nil
}
