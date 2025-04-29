package client

import "errors"

var (
	ErrNoKubeconfig = errors.New("no kubeconfig path provided and cannot infer from environment")
)
