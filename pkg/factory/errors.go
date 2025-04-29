package factory

import "errors"

var (
	ErrUnsupportedCloudProvider = errors.New("unsupported cloud provider")
	ErrFailedToCreateProvider   = errors.New("failed to create provider")
)
