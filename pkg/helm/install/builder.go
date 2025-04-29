package install

import "helm.sh/helm/v3/pkg/action"

// Builder ...
type Builder struct {
	install *action.Install
}

// NewBuilder ...
func NewBuilder(cfg *action.Configuration, opts ...Option) *Builder {
	install := action.NewInstall(cfg)

	for _, opt := range opts {
		opt(install)
	}

	return &Builder{install: install}
}

// Build ...
func (b *Builder) Build() *action.Install {
	return b.install
}
