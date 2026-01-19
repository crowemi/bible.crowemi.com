package bible_client

type Format string

const (
	FormatHtml Format = "html"
	FormatText Format = "text"
)

type Translation interface {
	Get(query string, config ...*TranslationConfig) (*TranslationResponse, error)
}

type TranslationResponse struct {
	Query     string `json:"query"`
	Canonical string `json:"canonical"`
	Passage   []byte `json:"passage"`
	Text      string `json:"text"`
}

type TranslationConfig struct {
	Format Format
}

// func GetBibleClient(config *config.Config, version string) Translation {
// 	var ret Translation
// 	// default to esv
// 	switch version {
// 	default:
// 		r := EsvClient{
// 			Token: config.EsvClientconfig.Token,
// 		}
// 		ret = &r
// 	}
// 	return ret
// }
