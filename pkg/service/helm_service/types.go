package helmservice

// InstallRequest ...
type InstallRequest struct {
	ReleaseName string
	ChartPath   string
	Values      map[string]any
}

// UpgradeRequest ...
type UpgradeRequest struct {
	ReleaseName string
	ChartPath   string
	Values      map[string]any
}

// UninstallRequest ...
type UninstallRequest struct {
	ReleaseName string
}
