package access

import "qdrant/pkg/provider"

// Info ...
type Info struct {
	RoleArn string
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
		RoleArn: "arn:aws:iam::account:role/eksRole",
	}, nil
}
