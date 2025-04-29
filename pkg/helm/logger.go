package helm

import (
	"qdrant/pkg/utils/logger"

	"github.com/rs/zerolog"
)

// Logger ...
type Logger interface {
	Debugf(format string, args ...any)
	Errorf(format string, args ...any)
	Infof(format string, args ...any)
}

// DefaultLogger ...
func DefaultLogger() Logger {
	return &zerologAdapter{logger.InitLogger()}
}

type zerologAdapter struct {
	logger zerolog.Logger
}

// Debugf ...
func (l *zerologAdapter) Debugf(format string, args ...any) {
	l.logger.Debug().Msgf(format, args...)
}

// Infof ...
func (l *zerologAdapter) Infof(format string, args ...any) {
	l.logger.Info().Msgf(format, args...)
}

// Errorf ...
func (l *zerologAdapter) Errorf(format string, args ...any) {
	l.logger.Error().Msgf(format, args...)
}
