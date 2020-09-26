package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	GomniauthKey      string
	GithubClientID    string
	GithubSecretValue string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		GomniauthKey:      cfg.Section("gomniauth").Key("security_key").String(),
		GithubClientID:    cfg.Section("github").Key("clientID").String(),
		GithubSecretValue: cfg.Section("github").Key("secret_value").String(),
	}
}
