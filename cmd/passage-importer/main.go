package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	firestore "github.com/crowemi-io/crowemi-go-utils/db/gcp"
	"github.com/crowemi/bible.crowemi.com/config"
	"github.com/crowemi/bible.crowemi.com/pkg/bible"
)

func main() {
	config, err := config.LoadConfig("../../.secret/config.json")
	if err != nil {
		fmt.Print(err)
	}
	log.Println(config.Crowemi.ClientName)

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
			Passage: record[0],
			Date:    record[1],
			Plan:    record[3],
			Summary: "",
			// "/page?q=Psalms.1-5"
			Link: fmt.Sprintf("/page?q=%s", strings.Replace(record[0], " ", ".", -1)),
		}
		passages = append(passages, p)
	}

	fs, err := config.FirestoreClient.Connect(context.TODO())
	if err != nil {
		fmt.Print(err)
	}
	defer fs.Close()

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
