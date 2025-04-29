package client

import (
	"path/filepath"

	"qdrant/pkg/helm"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"k8s.io/client-go/util/homedir"
)

// Client ...
type Client struct {
	config    *action.Configuration
	logger    helm.Logger
	namespace string
}

// NewClient ...
func NewClient(opts ...Option) (*Client, error) {
	options := &configOptions{}

	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	if options.kubeconfigPath == "" {
		if home := homedir.HomeDir(); home != "" {
			options.kubeconfigPath = filepath.Join(home, ".kube", "config")
		} else {
			return nil, ErrNoKubeconfig
		}
	}

	if options.namespace == "" {
		options.namespace = "default"
	}

	if options.logger == nil {
		options.logger = helm.DefaultLogger()
	}

	cfg := new(action.Configuration)

	if err := cfg.Init(
		kube.GetConfig(options.kubeconfigPath, "", options.namespace),
		options.namespace,
		"secrets",
		options.logger.Debugf,
	); err != nil {
		return nil, err
	}

	return &Client{
		config:    cfg,
		logger:    options.logger,
		namespace: options.namespace,
	}, nil
}

// Config ...
func (c *Client) Config() *action.Configuration {
	return c.config
}

// Namespace ....
func (c *Client) Namespace() string {
	return c.namespace
}

// Logger ...
func (c *Client) Logger() helm.Logger {
	return c.logger
}
