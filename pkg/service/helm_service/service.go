package helmservice

import (
	"qdrant/pkg/helm/client"
	"qdrant/pkg/helm/install"
	"qdrant/pkg/helm/uninstall"
	"qdrant/pkg/helm/upgrade"
)

var _ Interface = (*HelmService)(nil)

// HelmService ...
type HelmService struct {
	client *client.Client
}

// New ...
func New(c *client.Client) *HelmService {
	return &HelmService{
		client: c,
	}
}

// Install ...
func (s *HelmService) Install(req InstallRequest) error {
	if req.ReleaseName == "" || req.ChartPath == "" {
		return ErrInvalidRequest
	}
	return install.InstallReleaseDefaultV1(
		s.client,
		req.ReleaseName,
		req.ChartPath,
		req.Values,
	)
}

// Upgrade ...
func (s *HelmService) Upgrade(req UpgradeRequest) error {
	if req.ReleaseName == "" || req.ChartPath == "" {
		return ErrInvalidRequest
	}
	return upgrade.UpgradeReleaseDefaultV1(
		s.client,
		req.ReleaseName,
		req.ChartPath,
		req.Values,
	)
}

// Uninstall ...
func (s *HelmService) Uninstall(req UninstallRequest) error {
	if req.ReleaseName == "" {
		return ErrInvalidRequest
	}
	return uninstall.UninstallReleaseDefaultV1(
		s.client,
		req.ReleaseName,
	)
}
