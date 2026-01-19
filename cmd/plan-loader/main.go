package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/crowemi-io/crowemi-bible/config"
	"github.com/crowemi-io/crowemi-bible/pkg/bible"
	firestore "github.com/crowemi-io/crowemi-go-utils/db/gcp"
)

func main() {
	config, err := config.LoadConfig("../../.secret/config.json")
	if err != nil {
		fmt.Print(err)
	}
	log.Println(config.Crowemi.ClientName)

	fs, err := config.FirestoreClient.Connect(context.TODO())
	if err != nil {
		fmt.Print(err)
	}
	defer fs.Close()

	plans := []bible.Plan{
		{
			Name:        "Old Testament",
			Description: "Read the entire Old Testament (except Psalms) in a year.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "New Testament",
			Description: "Read the entire New Testament (except The Gospels) in a quarter.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, {
			Name:        "Psalms",
			Description: "Read the Psalms in a month.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, {
			Name:        "Gospels",
			Description: "Read all four Gospels in a month.",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	planXWalk := map[string]string{}
	for _, plan := range plans {
		ref, _, err := firestore.InsertOne(context.TODO(), fs, "Plan", plan)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("add: %s", ref.ID)
		planXWalk[plan.Name] = ref.ID
	}

	f, err := os.Open("dater.csv")
	if err != nil {
		log.Print(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	passages := []bible.PlanItem{}
	for idx, record := range records {
		if idx == 0 {
			continue
		}
		// if strings.Contains(record[0], ",") {
		// 	books := strings.Split(record[0], ",")
		// 	for _, book := range books {

		// 	}
		// }
		p := bible.PlanItem{
			PlanID:    planXWalk[record[2]],
			Summary:   "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		passages = append(passages, p)
	}

	var wg sync.WaitGroup

	for _, p := range passages {
		wg.Add(1)
		go func(passage bible.PlanItem) {
			defer wg.Done()
			ref, _, err := firestore.InsertOne(context.TODO(), fs, "PlanItem", passage)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("add: %s", ref.ID)
		}(p)
	}

	wg.Wait()

}
