package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type OAuth2Config struct {
	GoogleConfig   oauth2.Config
	FacebookConfig oauth2.Config
}

var AppOAuth2Config OAuth2Config

func InitOAuth2Configs() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %s", err)
	}

	fmt.Println("GOOGLE_REDIRECT_URL: ", os.Getenv("GOOGLE_REDIRECT_URL"))

	AppOAuth2Config.GoogleConfig = oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	AppOAuth2Config.FacebookConfig = oauth2.Config{
		RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URL"),
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: facebook.Endpoint,
	}
}
