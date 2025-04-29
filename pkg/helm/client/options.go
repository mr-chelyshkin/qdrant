package client

import "qdrant/pkg/helm"

// Option ...
type Option func(*configOptions) error

type configOptions struct {
	kubeconfigPath string
	namespace      string
	logger         helm.Logger
}

// WithKubeconfigPath ...
func WithKubeconfigPath(path string) Option {
	return func(opts *configOptions) error {
		opts.kubeconfigPath = path
		return nil
	}
}

// WithNamespace ...
func WithNamespace(namespace string) Option {
	return func(opts *configOptions) error {
		opts.namespace = namespace
		return nil
	}
}

// WithLogger ...
func WithLogger(l helm.Logger) Option {
	return func(opts *configOptions) error {
		opts.logger = l
		return nil
	}
}
