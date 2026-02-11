package main

import (
	"goauth/internal/config"
	"goauth/pkg/database"
	"goauth/pkg/logger"
	"log"
	"strings"
)

const APP_NAME = "goauth"

func main() {
	// Загрузка конфигурации
	cfg := config.GetConfig()

	// Логгер
	appLogger := logger.New(logger.AppLoggerLevelDebug, APP_NAME)

	// Подключение к БД, выполнение миграций
	db, err := database.New(database.Config{
		DSN:           "",
		MigrationPath: cfg.Db.MigrationPath,
		MaxConns:      int32(cfg.Db.MaxConns),
		MinConns:      int32(cfg.Db.MinConns),
	}, appLogger)

	if err != nil {
		log.Fatal("Ошибка подключения к БД/создание миграций: ", err)
	}
	defer db.Close()

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
