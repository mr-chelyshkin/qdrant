package install

import (
	"errors"

	"qdrant/pkg/helm/chart"
	"qdrant/pkg/helm/client"
)

// InstallReleaseDefaultV1 ...
func InstallReleaseDefaultV1(c *client.Client, releaseName, chartPath string, values map[string]any) error {
	return InstallRelease(
		c,
		releaseName,
		chartPath,
		values,
		WithNamespace(c.Namespace()),
		WithReleaseName(releaseName),
		WithAtomic(true),
		WithWait(true),
	)
}

// InstallRelease ...
func InstallRelease(c *client.Client, releaseName, chartPath string, values map[string]any, opts ...Option) error {
	builder := NewBuilder(c.Config(), opts...)
	install := builder.Build()

	ch, err := chart.Load(chartPath)
	if err != nil {
		return errors.Join(ErrFailedToLoadChart, err)
	}

	_, err = install.Run(ch, values)
	if err != nil {
		return errors.Join(ErrFailedToInstallRelease, err)
	}

	c.Logger().Infof("successfully installed release %s", releaseName)
	return nil
}
