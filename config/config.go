package config

import (
	"html/template"

	cfg "github.com/crowemi-io/crowemi-go-utils/config"
	t "github.com/crowemi/bible.crowemi.com/templates"
)

type Config struct {
	Templates *template.Template
	Esv       Esv         `json:"esv"`
	Crowemi   cfg.Crowemi `json:"crowemi_config"`
}

type Esv struct {
	Token string `json:"token"`
}

func LoadConfig(configPath string) (*Config, error) {
	appConfig, err := cfg.Bootstrap[Config](configPath)
	if err != nil {
		return appConfig, err
	}

	tmpl, err := t.GetTemplatesFS()
	if err != nil {
		return appConfig, err
	}
	appConfig.Templates = tmpl
	return appConfig, nil
}
