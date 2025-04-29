package install

import "errors"

var (
	ErrFailedToInstallRelease = errors.New("failed to install release")
	ErrFailedToLoadChart      = errors.New("failed to load chart for installation")
)
