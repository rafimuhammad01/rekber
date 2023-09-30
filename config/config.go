package config

import "time"

var (
	conf *Config
)

type Config struct {
	AppName                  string
	AccessTokenDuration      time.Duration
	RefreshTokenDuration     time.Duration
	JWTAccessTokenSecretKey  string
	JWTRefreshTokenSecretKey string
}

func Get() *Config {
	return conf
}

// Set is used for testing purpose only, do not use!
func Set(c Config) {
	conf = &c
}
