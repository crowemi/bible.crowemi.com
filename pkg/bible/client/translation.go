package translation

import (
	"html/template"

	"github.com/crowemi/bible.crowemi.com/config"
)

type Format string

const (
	FormatHtml Format = "html"
	FormatText Format = "text"
)

type Translation interface {
	GetPassage(query string, config ...*TranslationConfig) (*TranslationResponse, error)
}

type TranslationResponse struct {
	Query     string          `json:"query"`
	Canonical string          `json:"canonical"`
	Passages  []template.HTML `json:"passages"`
	Text      string          `json:"text"`
}

type TranslationConfig struct {
	Format Format
}

func GetBibleClient(config *config.Config, version string) Translation {
	var ret Translation
	// default to esv
	switch version {
	default:
		r := EsvClient{
			Token: config.Esv.Token,
		}
		ret = &r
	}
	return ret
}
