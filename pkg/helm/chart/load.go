package chart

import (
	"errors"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

// Load ...
func Load(path string) (*chart.Chart, error) {
	ch, err := loader.Load(path)
	if err != nil {
		return nil, errors.Join(ErrFailedToLoadChart, err)
	}
	return ch, nil
}
