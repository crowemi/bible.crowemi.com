package main

import (
	"net/http"
	"os"

	"github.com/crowemi-io/crowemi-bible/config"
	"github.com/crowemi-io/crowemi-bible/internal/handlers"
)

func main() {
	config, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		// TODO: log
		panic(err)
	}

	planHandler := handlers.PlanHandler{
		Config: config,
	}

	http.HandleFunc("/plan", planHandler.GetMany)
	http.ListenAndServe(":8001", nil)
}
