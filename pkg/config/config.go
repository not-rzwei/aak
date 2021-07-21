package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ScheduleId string   `env:"AAK_OG_SCHEDULE_ID,required"`
	ApiKey     string   `env:"AAK_OG_API_KEY,required"`
	SecretKey  string   `env:"AAK_SECRET_KEY,required"`
	CloseTags  []string `env:"AAK_CLOSE_TAGS"`
	CloseDelay int      `env:"AAK_CLOSE_DELAY" envDefault:"15"`
}

func GetConfig() *Config {
	configuration := &Config{}

	err := env.Parse(configuration)
	if err != nil {
		logrus.Error("Unable to parse environment variables: ", err)
	}

	return configuration
}
