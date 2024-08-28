package config

import (
	"fmt"
	"log/slog"
)

func ConfigureLogger(config *Config) error {
	switch config.Env().LogLevel {
	case "error":
		slog.SetLogLoggerLevel(slog.LevelError)

	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)

	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)

	default:
		return fmt.Errorf("undefined log level: %s", config.Env().LogLevel)
	}

	return nil
}
