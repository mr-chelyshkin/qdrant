package helmservice

type Interface interface {
	Install(req InstallRequest) error
	Upgrade(req UpgradeRequest) error
	Uninstall(req UninstallRequest) error
}
