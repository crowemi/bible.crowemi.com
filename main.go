package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/crowemi-io/crowemi-go-utils/db"
	firestore "github.com/crowemi-io/crowemi-go-utils/db/gcp"
	"github.com/crowemi/bible.crowemi.com/config"
	"github.com/crowemi/bible.crowemi.com/pkg/bible"
	client "github.com/crowemi/bible.crowemi.com/pkg/bible/client"
)

const (
	Collection = "passages"
)

type data struct {
	Scripture []bible.PlanItem
}

func main() {
	config, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		// TODO: log
		panic(err)
	}

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		query := r.URL.Query()
		passageQuery := query.Get("q")
		version := query.Get("version")

		bibleClient := client.GetBibleClient(config, version)
		passage, err := bibleClient.GetPassage(passageQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		page := config.Templates.Lookup("page")
		err = page.Execute(w, passage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs, err := config.FirestoreClient.Connect(context.TODO())
		if err != nil {
			// TODO: log
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer fs.Close()

		today := time.Now()
		filters := []db.Filter{}
		filter := db.Filter{
			Field:    "date",
			Operator: "==",
			Value:    today.Format("January 2, 2006"),
		}
		filters = append(filters, filter)

		passages, err := firestore.GetMany[bible.PlanItem](context.TODO(), fs, Collection, filters)
		if err != nil {
			// TODO: log
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bibleClient := client.GetBibleClient(config, "esv")

		for idx, passage := range *passages {
			if passage.Summary != "" {
				continue
			}
			query := strings.Replace(passage.Passage, " ", ".", -1)
			passageResponse, err := bibleClient.GetPassage(query, &client.TranslationConfig{
				Format: client.FormatText,
			})
			if err != nil {
				// TODO: log
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			summary, err := config.GenAI.Generate(context.TODO(), fmt.Sprintf("Summarize this passage in light of the Gospel of Jesus Christ in 1-3 sentences, don't use the phrase \"In light of the Gospel\", and no fluff: %s", string(passageResponse.Passages[0])))
			if err != nil {
				// TODO: log
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			passage.Summary = summary
			(*passages)[idx] = passage
			// write summary to firestore
			// _, err = firestore.UpdateOne(context.TODO(), fs, Collection, passage.PlanItemID, update)
			// if err != nil {
			// 	// TODO: log
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
		}

		data := data{
			Scripture: *passages,
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page := config.Templates.Lookup("home")
		err = page.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
