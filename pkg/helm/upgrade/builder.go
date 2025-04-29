package upgrade

import "helm.sh/helm/v3/pkg/action"

// Builder ...
type Builder struct {
	upgrade *action.Upgrade
}

// NewBuilder ...
func NewBuilder(cfg *action.Configuration, opts ...Option) *Builder {
	upgrade := action.NewUpgrade(cfg)

	for _, opt := range opts {
		opt(upgrade)
	}

	return &Builder{upgrade: upgrade}
}

// Build ...
func (b *Builder) Build() *action.Upgrade {
	return b.upgrade
}
