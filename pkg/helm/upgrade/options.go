package upgrade

import (
	"time"

	"helm.sh/helm/v3/pkg/action"
)

// Option ...
type Option func(*action.Upgrade)

// WithNamespace ...
func WithNamespace(namespace string) Option {
	return func(u *action.Upgrade) {
		u.Namespace = namespace
	}
}

// WithAtomic ...
func WithAtomic(atomic bool) Option {
	return func(u *action.Upgrade) {
		u.Atomic = atomic
	}
}

// WithWait ...
func WithWait(wait bool) Option {
	return func(u *action.Upgrade) {
		u.Wait = wait
	}
}

// WithTimeout ...
func WithTimeout(timeoutSeconds int64) Option {
	return func(u *action.Upgrade) {
		u.Timeout = time.Duration(timeoutSeconds) * time.Second
	}
}
