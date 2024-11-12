package config

import (
	"os"
	"sync"
)

var (
	once      sync.Once
	appConfig *Config
)

// Config struct holds the environment variables
type Config struct {
	Port                 string
	NotificationURL      string
	ReturnSuccessURL     string
	ReturnFailureURL     string
	ReturnCancelURL      string
	ReturnPendingURL     string
	TerminalId           string
	EmerchantPayUsername string
	EmerchantPayPassword string

	// Emerchantpay api urls
	EmerchantPayWpfURL string
}

// GetConfig returns the singleton instance of the configuration
func GetConfig() *Config {
	once.Do(func() {
		appConfig = &Config{
			Port:                 os.Getenv("APP_PORT"),
			NotificationURL:      os.Getenv("WEBHOOK_NOTIFICATION_URL"),
			ReturnSuccessURL:     os.Getenv("RETURN_SUCCESS_URL"),
			ReturnFailureURL:     os.Getenv("RETURN_FAILURE_URL"),
			ReturnCancelURL:      os.Getenv("RETURN_CANCEL_URL"),
			ReturnPendingURL:     os.Getenv("RETURN_PENDING_URL"),
			TerminalId:           os.Getenv("EMERCHANTPAY_TERMINAL_ID"),
			EmerchantPayUsername: os.Getenv("EMERCHANTPAY_USERNAME"),
			EmerchantPayPassword: os.Getenv("EMERCHANTPAY_PASSWORD"),
			EmerchantPayWpfURL:   os.Getenv("EMERCHANTPAY_WPF_API_URL"),
		}
	})
	return appConfig
}
