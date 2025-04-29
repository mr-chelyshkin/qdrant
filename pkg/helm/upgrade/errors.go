package upgrade

import "errors"

var (
	ErrFailedToUpgradeRelease = errors.New("failed to upgrade release")
	ErrFailedToLoadChart      = errors.New("failed to load chart for upgrade")
)
