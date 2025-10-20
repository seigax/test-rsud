package config

import (
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
)

func NewSMTPClient() lib.SMTPClient {
	return lib.SMTPClient{
		Host:     viper.GetString("SMTP_HOST"),
		Port:     viper.GetString("SMTP_PORT"),
		Password: viper.GetString("SMTP_PASSWORD"),
		Username: viper.GetString("SMTP_USERNAME"),
		Sender:   viper.GetString("SMTP_SENDER"),
	}
}
