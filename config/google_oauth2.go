package config

import (
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOauth2() *lib.GoogleOauth2 {
	return &lib.GoogleOauth2{
		&oauth2.Config{
			ClientID:     viper.GetString("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: viper.GetString("GOOGLE_OAUTH_CLIENT_SECRET"),
			RedirectURL:  viper.GetString("GOOGLE_OAUTH_CALLBACK"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}
