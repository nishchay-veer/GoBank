package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `default:"postgres" mapstructure:"DB_DRIVER"`
	DBSource             string        `default:"postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string        `default:":8080" mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `default:":5000" mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `default:"secret" mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `default:"1h" mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `default:"24h" mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisAddress         string        `default:":6379" mapstructure:"REDIS_ADDRESS"`
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
