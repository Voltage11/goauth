package main

import (
	"goauth/internal/config"
	"goauth/pkg/logger"
	"strings"
)

const APP_NAME = "goauth"

func main() {
	// Загрузка конфигурации
	cfg := config.GetConfig()

	// Логгер
	appLogger := logger.New(logger.AppLoggerLevelDebug, APP_NAME)

	// До запуска сервера установим уровень логирования из настроек
	appLogger.SetLevel(getLogLevelFromStr(cfg.Logger.Level))
}

func getLogLevelFromStr(logLevelStr string) logger.AppLoggerLevel {
	switch strings.ToLower(logLevelStr) {
	case "debug":
		return logger.AppLoggerLevelDebug
	case "info":
		return logger.AppLoggerLevelInfo
	case "warn":
		return logger.AppLoggerLevelWarn
	case "error":
		return logger.AppLoggerLevelError
	default:
		return logger.AppLoggerLevelInfo
	}
}
