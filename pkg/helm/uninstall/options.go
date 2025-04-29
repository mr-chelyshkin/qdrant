package uninstall

import "helm.sh/helm/v3/pkg/action"

// Option ...
type Option func(*action.Uninstall)

// WithKeepHistory ...
func WithKeepHistory(keep bool) Option {
	return func(u *action.Uninstall) {
		u.KeepHistory = keep
	}
}
