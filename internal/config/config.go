package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type DbConfig struct {
	Host          string `env:"DB_HOST" env-default:"localhost"`
	Port          string `env:"DB_PORT" env-default:"5432"`
	Name          string `env:"DB_NAME" env-default:"postgres"`
	User          string `env:"DB_USER" env-required:"true"`
	Password      string `env:"DB_PASSWORD" env-required:"true"`
	MigrationPath string `env:"DB_MIGRATION_PATH" env-required:"true"`
	MaxConns      int    `env:"DB_MAX_CONNS" env-required:"25"`
	MinConns      int    `env:"DB_MIN_CONNS" env-required:"5"`
}

type SecretConfig struct {
	Jwt  string `env:"SECRET_JWT" env-required:"true"`
	Hash string `env:"SECRET_HASH" env-required:"true"`
}

type ServerConfig struct {
	Host string   `env:"SERVER_HOST" env-required:"true"`
	Port int      `env:"SERVER_PORT" env-required:"true"`
	Cors []string `env:"SERVER_CORS" env-separator:","`
}

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL" env-default:"debug"`
}

type SmtpConfig struct {
	Host     string `env:"EMAIL_SMTP_HOST" env-required:"true"`
	Port     int    `env:"EMAIL_SMTP_PORT" env-required:"true"`
	User     string `env:"EMAIL_SMTP_USERNAME" env-required:"true"`
	Password string `env:"EMAIL_SMTP_PASSWORD" env-required:"true"`
}

type AppConfig struct {
	Db     DbConfig
	Secret SecretConfig
	Server ServerConfig
	Logger LoggerConfig
	Smtp   SmtpConfig
}

func GetConfig() *AppConfig {
	var cfg AppConfig

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatal("Ошибка загрузки переменных окружения, ", err)
	}

	return &cfg
}
