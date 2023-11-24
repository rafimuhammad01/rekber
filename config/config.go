package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	conf *Config
)

type (
	Config struct {
		App      AppConfig  `mapstructure:"app"`
		JWT      JWTConfig  `mapstructure:"jwt"`
		PSQL     PSQLConfig `mapstructure:"psql"`
		Firebase Firebase   `mapstructure:"firebase"`
	}

	AppConfig struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
	}

	JWTConfig struct {
		AccessToken  TokenConfig `mapstructure:"access_token"`
		RefreshToken TokenConfig `mapstructure:"refresh_token"`
	}

	TokenConfig struct {
		Duration  time.Duration `maspstructure:"duration"`
		SecretKey string        `maspstructure:"secret_key"`
	}

	PSQLConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		UserName string `mapstructure:"user_name"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
		SSLMode  string `mapstructure:"ssl_mode"`
	}

	Firebase struct {
		APIKey  string `mapstructure:"api_key"`
		AuthURL string `mapstructure:"url"`
	}
)

func Get() *Config {
	return conf
}

// Set is used for testing purpose only, do not use!
func Set(c Config) {
	conf = &c
}

func SetFromFile(fileName string) {
	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatalf("fatal error config file: %v", err.Error())
	}

	var configFromViper Config
	err = viper.Unmarshal(&configFromViper)
	if err != nil {
		log.Fatalf("unable to decode into config struct: %v", err.Error())
	}

	conf = &configFromViper
}
