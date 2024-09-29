package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `default:"postgres" mapstructure:"DB_DRIVER"`
	DBSource string `default:"postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" mapstructure:"DB_SOURCE"`
	ServerAddress string `default:":8080" mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `default:"secret" mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `default:"1h" mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return 

}