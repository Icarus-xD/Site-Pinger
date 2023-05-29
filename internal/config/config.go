package config

import "github.com/spf13/viper"

type Config struct {
	AppPort     string `mapstructure:"APP_PORT"`
	PostgresUrl string `mapstructure:"POSTGRES_URL"`
	RedisAddr   string `mapstructure:"REDIS_ADDR"`
	ClickhouseAddr   string `mapstructure:"CLICKHOUSE_ADDR"`
}

func GetConfig() (*Config, error) {
	cfg := new(Config)

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}