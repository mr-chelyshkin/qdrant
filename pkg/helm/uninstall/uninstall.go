package uninstall

import (
	"errors"
	"qdrant/pkg/helm/client"

	"helm.sh/helm/v3/pkg/action"
)

// UninstallReleaseDefaultV1 ...
func UninstallReleaseDefaultV1(c *client.Client, releaseName string) error {
	return UninstallRelease(c, releaseName, WithKeepHistory(false))
}

// UninstallRelease ...
func UninstallRelease(c *client.Client, releaseName string, opts ...Option) error {
	uninstall := action.NewUninstall(c.Config())

	for _, opt := range opts {
		opt(uninstall)
	}

	_, err := uninstall.Run(releaseName)
	if err != nil {
		return errors.Join(ErrFailedToUninstallRelease, err)
	}

	c.Logger().Infof("successfully uninstalled release %s", releaseName)
	return nil
}
