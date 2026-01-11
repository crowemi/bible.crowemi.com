package config

import (
	"html/template"

	genai "github.com/crowemi-io/crowemi-go-utils/ai/gcp"
	"github.com/crowemi-io/crowemi-go-utils/config"
	firestore "github.com/crowemi-io/crowemi-go-utils/db/gcp"
	"github.com/crowemi/bible.crowemi.com/templates"
)

type Config struct {
	Templates       *template.Template
	Esv             Esv                `json:"esv"`
	BibleApi        BibleApi           `json:"bible_api"`
	Crowemi         config.Crowemi     `json:"crowemi_config"`
	GoogleCloud     config.GoogleCloud `json:"google_cloud"`
	FirestoreClient firestore.Client
	GenAI           genai.Client
}

type Esv struct {
	Token string `json:"token"`
}

type BibleApi struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func LoadConfig(configPath string) (*Config, error) {
	appConfig, err := config.Bootstrap[Config](configPath)
	if err != nil {
		return appConfig, err
	}

	tmpl, err := templates.GetTemplatesFS()
	if err != nil {
		return appConfig, err
	}
	appConfig.Templates = tmpl
	appConfig.FirestoreClient = firestore.Client{
		Config: &appConfig.GoogleCloud,
	}
	appConfig.GenAI = genai.Client{
		Config: &appConfig.GoogleCloud,
	}
	return appConfig, nil
}
