package install

import "helm.sh/helm/v3/pkg/action"

// Option ...
type Option func(*action.Install)

// WithNamespace ...
func WithNamespace(namespace string) Option {
	return func(i *action.Install) {
		i.Namespace = namespace
	}
}

// WithReleaseName ...
func WithReleaseName(name string) Option {
	return func(i *action.Install) {
		i.ReleaseName = name
	}
}

// WithWait ...
func WithWait(wait bool) Option {
	return func(i *action.Install) {
		i.Wait = wait
	}
}

// WithAtomic ...
func WithAtomic(atomic bool) Option {
	return func(i *action.Install) {
		i.Atomic = atomic
	}
}
