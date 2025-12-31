package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/crowemi/bible.crowemi.com/config"
)

type data struct {
	Scripture []passages
}

type passages struct {
	Passage string
	Date    string
	Plan    string
	Summary string
	Link    string
}

type esvResponse struct {
	Query     string          `json:"query"`
	Canonical string          `json:"canonical"`
	Passages  []template.HTML `json:"passages"`
}

func main() {
	config, err := config.LoadConfig("/Users/crowemi/code/bible.crowemi.com/.secret/config.json")
	if err != nil {
		panic(err)
	}

	data := data{
		Scripture: []passages{
			{
				Passage: "Genesis 1-2",
				Date:    "January 1, 2026",
				Plan:    "Old Testament",
				Link:    "/page?q=Genesis.1-2",
				Summary: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			},
			{
				Passage: "Psalms 1-5",
				Date:    "January 1, 2026",
				Plan:    "Psalms",
				Link:    "/page?q=Psalms.1-5",
				Summary: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			},
			{
				Passage: "Acts 1-2",
				Date:    "January 1, 2026",
				Plan:    "New Testament",
				Link:    "/page?q=Acts.1-2",
				Summary: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			},
			{
				Passage: "Matthew 1-3",
				Date:    "January 1, 2026",
				Plan:    "Gospels",
				Link:    "/page?q=Matthew.1-3",
				Summary: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			},
		},
	}

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		query := r.URL.Query()
		esvQuery := query.Get("q")

		token := "Token " + config.Esv.Token
		// 'inline-styles': false,  // Don't include inline styles
		// 'include-css-link': false,  // Don't include their CSS
		// 'wrapping-div': true,  // Wrap in a div
		// 'div-classes': 'esv-passage',  // Add your custom class
		// 'paragraph-tag': 'p'  // Use standard p tags
		uri := fmt.Sprintf("https://api.esv.org/v3/passage/html/?q=%s&inline-styles=false&include-css-link=false&wrapping-div=true&div-classes=esv-passage&paragraph-tag=p&include-book-titles=false&include-footnotes=false", esvQuery)

		req, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		req.Header.Add("Authorization", token)
		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer resp.Body.Close()

		esvResp := &esvResponse{}
		err = json.Unmarshal(body, esvResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		page := config.Templates.Lookup("page")
		err = page.Execute(w, esvResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page := config.Templates.Lookup("home")
		err = page.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
