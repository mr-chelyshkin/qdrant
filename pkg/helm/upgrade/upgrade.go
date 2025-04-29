package upgrade

import (
	"errors"

	"qdrant/pkg/helm/chart"
	"qdrant/pkg/helm/client"
)

// UpgradeReleaseDefaultV1 ...
func UpgradeReleaseDefaultV1(c *client.Client, releaseName, chartPath string, values map[string]any) error {
	return UpgradeRelease(
		c,
		releaseName,
		chartPath,
		values,
		WithNamespace(c.Namespace()),
		WithAtomic(true),
		WithWait(true),
	)
}

// UpgradeRelease ...
func UpgradeRelease(c *client.Client, releaseName, chartPath string, values map[string]any, opts ...Option) error {
	builder := NewBuilder(c.Config(), opts...)
	upgrade := builder.Build()

	ch, err := chart.Load(chartPath)
	if err != nil {
		return errors.Join(ErrFailedToLoadChart, err)
	}

	_, err = upgrade.Run(releaseName, ch, values)
	if err != nil {
		return errors.Join(ErrFailedToUpgradeRelease, err)
	}

	c.Logger().Infof("successfully upgraded release %s", releaseName)
	return nil
}
