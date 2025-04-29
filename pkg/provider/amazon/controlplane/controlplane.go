package controlplane

import (
	"context"
	"errors"

	"qdrant/pkg/provider"
	"qdrant/pkg/provider/amazon/inputs"

	"github.com/aws/aws-sdk-go-v2/aws"
	sdkConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

type Info struct {
	ID         string
	Name       string
	Region     string
	Status     string
	Endpoint   string
	Kubeconfig any
}

// Service ...
type Service struct {
	client    *eks.Client
	region    string
	roleArn   string
	subnetIds []string
}

// New ...
func New(cfg inputs.ControlPlaneInputConfig) (*Service, error) {
	awsCfg, err := sdkConfig.LoadDefaultConfig(context.Background(), sdkConfig.WithRegion(cfg.Region))
	if err != nil {
		return nil, errors.Join(ErrFailedToLoadAWSConfig, err)
	}

	return &Service{
		client:    eks.NewFromConfig(awsCfg),
		region:    cfg.Region,
		roleArn:   cfg.RoleArn,
		subnetIds: cfg.SubnetIds,
	}, nil
}

// Create ...
func (s *Service) Create(cfg provider.ClusterConfig) (*Info, error) {
	input := &eks.CreateClusterInput{
		Name:    aws.String(cfg.Name),
		Version: aws.String(cfg.K8sVersion),
		RoleArn: aws.String(s.roleArn),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			SubnetIds: s.subnetIds,
		},
	}

	out, err := s.client.CreateCluster(context.Background(), input)
	if err != nil {
		return nil, errors.Join(ErrFailedToCreateCluster, err)
	}

	return &Info{
		ID:       aws.ToString(out.Cluster.Arn),
		Name:     aws.ToString(out.Cluster.Name),
		Region:   s.region,
		Status:   string(out.Cluster.Status),
		Endpoint: aws.ToString(out.Cluster.Endpoint),
	}, nil
}

// DeleteCluster ...
func (s *Service) DeleteCluster(id string) error {
	_, err := s.client.DeleteCluster(context.Background(), &eks.DeleteClusterInput{
		Name: aws.String(id),
	})
	if err != nil {
		return errors.Join(ErrFailedToDeleteCluster, err)
	}
	return nil
}

// GetClusterInfo ...
func (s *Service) GetClusterInfo(id string) (*Info, error) {
	out, err := s.client.DescribeCluster(context.Background(), &eks.DescribeClusterInput{
		Name: aws.String(id),
	})
	if err != nil {
		return nil, errors.Join(ErrFailedToDescribeCluster, err)
	}

	return &Info{
		Endpoint:   aws.ToString(out.Cluster.Endpoint),
		Name:       aws.ToString(out.Cluster.Name),
		ID:         aws.ToString(out.Cluster.Arn),
		Status:     string(out.Cluster.Status),
		Region:     s.region,
		Kubeconfig: nil,
	}, nil
}
