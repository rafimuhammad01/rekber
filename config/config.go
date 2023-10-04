package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	conf *Config
)

type Config struct {
	AppName string `mapstructure:"app_name"`
	Port    string `mapstructure:"port"`

	// JWT
	AccessTokenDuration      time.Duration `mapstructure:"access_token_duration"`
	RefreshTokenDuration     time.Duration `mapstructure:"refresh_token_duration"`
	JWTAccessTokenSecretKey  string        `mapstructure:"jwt_access_token_secret_key"`
	JWTRefreshTokenSecretKey string        `mapstructure:"jwt_refresh_token_secret_key"`

	// PSQL
	PSQLHost     string `mapstructure:"psql_host"`
	PSQLPort     string `mapstructure:"psql_port"`
	PSQLUserName string `mapstructure:"psql_username"`
	PSQLDBName   string `mapstructure:"psql_db_name"`
	PSQLPassword string `mapstructure:"psql_password"`
	PSQLSSLMode  string `mapstructure:"psql_ssl_mode"`
}

func Get() *Config {
	return conf
}

// Set is used for testing purpose only, do not use!
func Set(c Config) {
	conf = &c
}

func SetFromFile(fileName string) {
	viper.SetConfigName(fileName)   // name of config file (without extension)
	viper.SetConfigType("json")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in

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
