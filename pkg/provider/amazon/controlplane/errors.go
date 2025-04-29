package controlplane

import "errors"

var (
	ErrFailedToCreateCluster   = errors.New("failed to create EKS cluster")
	ErrFailedToDeleteCluster   = errors.New("failed to delete EKS cluster")
	ErrFailedToDescribeCluster = errors.New("failed to describe EKS cluster")
	ErrFailedToLoadAWSConfig   = errors.New("failed to load AWS config")
)
