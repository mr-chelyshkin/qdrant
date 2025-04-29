package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const (
	EnvLogLevel = "LOG_LEVEL"
)

// InitLogger ...
func InitLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "",
		PartsOrder: []string{
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		},
	}
	return zerolog.New(output).Level(getLogLevel()).With().Logger()
}

func getLogLevel() zerolog.Level {
	switch strings.ToLower(os.Getenv(EnvLogLevel)) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn", "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
