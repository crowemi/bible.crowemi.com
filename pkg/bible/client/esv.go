package bible_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EsvClient struct {
	Token    string `json:"token"`
	Endpoint string `json:"endpoint"`
}

type EsvResponse struct {
	Query       string   `json:"query"`
	Canonical   string   `json:"canonical"`
	Parsed      [][]int  `json:"parsed"`
	Passages    []string `json:"passages"`
	PassageMeta []struct {
		Canonical    string `json:"canonical"`
		ChapterStart []int  `json:"chapter_start"`
		ChapterEnd   []int  `json:"chapter_end"`
		PrevVerse    int    `json:"prev_verse"`
		NextVerse    int    `json:"next_verse"`
		PrevChapter  []int  `json:"prev_chapter"`
		NextChapter  []int  `json:"next_chapter"`
	} `json:"passage_meta"`
}

func (e *EsvClient) GetPassage(query string, opts ...map[string]string) (*EsvResponse, error) {
	token := "Token " + e.Token
	uri := fmt.Sprintf("%s/passage/text/?q=%s", e.Endpoint, query)
	for _, opt := range opts {
		for k, v := range opt {
			uri += fmt.Sprintf("&%s=%s", k, v)
		}
	}

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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ESV status not okay: %s", resp.Status)
	}

	// maybe we should read this in chunks
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	esvResp := &EsvResponse{}
	err = json.Unmarshal(body, esvResp)
	if err != nil {
		return nil, err
	}

	return esvResp, nil
}
