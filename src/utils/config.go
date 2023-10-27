package utils

import "github.com/spf13/viper"

type Config struct {
	SMTP_HOST      string
	SMTP_PORT      string
	SMTP_USERNAME  string
	SMTP_PASSWORD  string
	SMTP_FROM      string
	CLICKATELL_KEY string
}

var envConfig *viper.Viper
var Env *Config

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Env = &config
	return
}
