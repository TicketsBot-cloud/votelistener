package config

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	SentryDsn *string       `env:"SENTRY_DSN"`
	JsonLogs  bool          `env:"JSON_LOGS" envDefault:"false"`
	LogLevel  zapcore.Level `env:"LOG_LEVEL" envDefault:"info"`

	ServerAddr  string `env:"SERVER_ADDR" envDefault:":8080"`
	DatabaseUri string `env:"DATABASE_URI,required"`

	DblToken string `env:"DBL_TOKEN,required"`
}

func LoadFromEnv() (Config, error) {
	return env.ParseAs[Config]()
}
