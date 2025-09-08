package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	RateLimitTTLDuration time.Duration
	RateLimitMaxAttempts uint
	OTPTTLDuration       time.Duration
	Port                 uint16
	DBConnection         string
	JwtSecret            string
	JwtTTLDuration       time.Duration
}

func ReadConfig() (*AppConfig, error) {
	RateLimitTTLDuration, err := time.ParseDuration(os.Getenv("RATE_LIMIT_TTL_DURATION"))
	if err != nil {
		RateLimitTTLDuration = time.Minute * 10
	}

	RateLimitMaxAttempts, err := strconv.ParseUint(os.Getenv("RATE_LIMIT_MAX_ATTEMPTS"), 10, 64)
	if err != nil {
		RateLimitMaxAttempts = 3
	}

	OTPTTLDuration, err := time.ParseDuration(os.Getenv("OTP_TTL_DURATION"))
	if err != nil {
		OTPTTLDuration = time.Minute * 2
	}

	Port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)
	if err != nil {
		Port = 3000
	}

	DBConnection := os.Getenv("DB_CONNECTION")
	if DBConnection == "" {
		return nil, errors.New("You Must Provide 'DB_CONNECTION' environment variable")
	}

	JwtSecret := os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		return nil, errors.New("You Must Provide 'JWT_SECRET' environment variable")
	}

	JwtTTLDuration, err := time.ParseDuration(os.Getenv("JWT_TTL_DURATION"))
	if err != nil {
		JwtTTLDuration = time.Hour * 24
	}

	config := AppConfig{
		RateLimitTTLDuration: RateLimitTTLDuration,
		RateLimitMaxAttempts: uint(RateLimitMaxAttempts),
		OTPTTLDuration:       OTPTTLDuration,
		Port:                 uint16(Port),
		DBConnection:         DBConnection,
		JwtSecret:            JwtSecret,
		JwtTTLDuration:       JwtTTLDuration,
	}

	return &config, nil
}
