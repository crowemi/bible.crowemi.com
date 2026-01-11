package translation

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type EsvClient struct {
	Token string `json:"token"`
}

type EsvResponse struct {
	Query     string          `json:"query"`
	Canonical string          `json:"canonical"`
	Passages  []template.HTML `json:"passages"`
}

func (e *EsvClient) GetPassage(query string, config ...*TranslationConfig) (*TranslationResponse, error) {
	var ret *TranslationResponse
	if config != nil {
		switch config[0].Format {
		case FormatText:
			text, err := e.GetPassageText(query)
			if err != nil {
				return nil, err
			}
			ret = text
		case FormatHtml:
			html, err := e.GetPassageHtml(query)
			if err != nil {
				return nil, err
			}
			ret = html
		default:
			return nil, fmt.Errorf("unknown format: %s", config[0].Format)
		}
	} else {
		// default GetPassage
		html, err := e.GetPassageHtml(query)
		if err != nil {
			return nil, err
		}
		ret = html
	}
	return ret, nil
}
func (e *EsvClient) GetPassageText(query string) (*TranslationResponse, error) {
	token := "Token " + e.Token
	uri := fmt.Sprintf("https://api.esv.org/v3/passage/text/?q=%s&include-book-titles=false&include-footnotes=false", query)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// maybe we should read this in chunks
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	translationResp := &TranslationResponse{}
	err = json.Unmarshal(body, translationResp)
	if err != nil {
		return nil, err
	}

	return translationResp, nil
}
func (e *EsvClient) GetPassageHtml(query string) (*TranslationResponse, error) {

	token := "Token " + e.Token
	uri := fmt.Sprintf("https://api.esv.org/v3/passage/html/?q=%s&inline-styles=false&include-css-link=false&wrapping-div=true&div-classes=esv-passage&paragraph-tag=p&include-book-titles=false&include-footnotes=false", query)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// maybe we should read this in chunks
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	translationResp := &TranslationResponse{}
	err = json.Unmarshal(body, translationResp)
	if err != nil {
		return nil, err
	}

	return translationResp, nil
}
