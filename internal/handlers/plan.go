package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/crowemi-io/crowemi-bible/config"
	"github.com/crowemi-io/crowemi-bible/pkg/bible"
	"github.com/crowemi-io/crowemi-go-utils/db"
	firestore "github.com/crowemi-io/crowemi-go-utils/db/gcp"
)

const Collection = "passages"

type PlanHandler struct {
	Config *config.Config
}

func (h *PlanHandler) GetMany(w http.ResponseWriter, r *http.Request) {
	fs, err := h.Config.FirestoreClient.Connect(context.TODO())
	if err != nil {
		// TODO: log
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer fs.Close()

	filters := []db.Filter{}
	filter := db.Filter{
		Field:    "date",
		Operator: "==",
		Value:    time.Now().Format("January 2, 2006"),
	}
	filters = append(filters, filter)

	planItems, err := firestore.GetMany[bible.PlanItem](context.TODO(), fs, Collection, filters)
	if err != nil {
		// TODO: log
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ret, err := json.Marshal(&planItems)
	if err != nil {
		// TODO: log
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(ret)
}
