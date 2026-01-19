package config

import (
	genai "github.com/crowemi-io/crowemi-go-utils/ai/gcp"
	"github.com/crowemi-io/crowemi-go-utils/config"
	firestore "github.com/crowemi-io/crowemi-go-utils/db/gcp"
)

type Config struct {
	EsvClientconfig EsvClientConfig    `json:"esv_client"`
	BibleApi        BibleApi           `json:"bible_api"`
	Crowemi         config.Crowemi     `json:"crowemi_config"`
	GoogleCloud     config.GoogleCloud `json:"google_cloud"`
	FirestoreClient firestore.Client
	GenAI           genai.Client
}

type EsvClientConfig struct {
	Token    string `json:"token"`
	Endpoint string `json:"endpoint"`
}

type BibleApi struct {
	Key      string `json:"key"`
	Endpoint string `json:"endpoint"`
}

func LoadConfig(configPath string) (*Config, error) {
	appConfig, err := config.Bootstrap[Config](configPath)
	if err != nil {
		return appConfig, err
	}
	appConfig.FirestoreClient = firestore.Client{
		Config: &appConfig.GoogleCloud,
	}
	appConfig.GenAI = genai.Client{
		Config: &appConfig.GoogleCloud,
	}
	return appConfig, nil
}
