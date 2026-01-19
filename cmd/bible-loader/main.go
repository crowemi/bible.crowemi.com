package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/crowemi-io/crowemi-bible/config"
	"github.com/crowemi-io/crowemi-bible/pkg/bible"
	"github.com/crowemi-io/crowemi-go-utils/db"
	f "github.com/crowemi-io/crowemi-go-utils/db/gcp"
)

type BibleLoader struct {
	Config          *config.Config
	FirestoreClient *firestore.Client
}

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fs, err := config.FirestoreClient.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()

	loader := BibleLoader{
		Config:          config,
		FirestoreClient: fs,
	}

	// loader.loadBooks()
	loader.loadChapters()

}

func (bl *BibleLoader) loadBooks() {
	resourcesPath := os.Getenv("RESOURCES_PATH")
	file, err := os.Open(fmt.Sprintf("%s/Books.csv", resourcesPath))
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for idx, r := range records {
		if idx == 0 {
			// skip header
			continue
		}
		chapterCount, _ := strconv.Atoi(r[10])
		verseCount, _ := strconv.Atoi(r[11])
		peopleCount, _ := strconv.Atoi(r[12])
		placeCount, _ := strconv.Atoi(r[13])
		book := bible.Book{
			ID:           r[0],
			Name:         r[2],
			Osis:         r[0],
			ShortName:    r[5],
			Number:       r[1],
			BookDivision: r[3],
			Testament:    r[4],
			Slug:         r[6],
			YearWritten:  r[7],
			PlaceWritten: r[8],
			ChapterCount: chapterCount, // chapterCount,
			VerseCount:   verseCount,   // verseCount,
			Writers:      r[11],        // writers,
			PeopleCount:  peopleCount,  // peopleCount,
			PlaceCount:   placeCount,   // placeCount,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		_, _, err := f.InsertOne(context.TODO(), bl.FirestoreClient, "bible", book)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(book)
	}

}

func (bl *BibleLoader) loadChapters() {
	resourcesPath := os.Getenv("RESOURCES_PATH")
	file, err := os.Open(fmt.Sprintf("%s/Chapters.csv", resourcesPath))
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	idx := 0
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		if idx == 0 {
			// skip header
			idx++
			continue
		}
		peopleCount, _ := strconv.Atoi(row[6])
		placeCount, _ := strconv.Atoi(row[7])
		writerCount, _ := strconv.Atoi(row[9])

		chapter := bible.Chapter{
			ID:          strings.ToLower(row[0]),
			Osis:        row[0],
			Number:      row[2], // chapterNum,
			Verses:      []bible.Verse{},
			Summary:     "",
			Writer:      row[3],
			Slug:        row[5],
			PeopleCount: peopleCount,
			PlaceCount:  placeCount,
			WriterCount: writerCount,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		filter := db.Filter{
			Field:    "osis",
			Operator: "==",
			Value:    row[1],
		}
		filters := []db.Filter{}
		filters = append(filters, filter)

		book, id, err := f.GetOne[bible.Book](context.TODO(), bl.FirestoreClient, "bible", filters)
		if err != nil {
			log.Fatal(err)
		}

		book.Chapters = append(book.Chapters, chapter)

		updates := []f.Update{}
		updates = append(updates, f.Update{
			Path:  "chapters",
			Value: book.Chapters,
		})
		_, err = f.UpdateOne(context.TODO(), bl.FirestoreClient, "bible", id, updates)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (bl *BibleLoader) loadVerses() {}
