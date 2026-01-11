package bibleloader

import (
	"log"

	"github.com/crowemi/bible.crowemi.com/config"
)

func main() {
	config, err := config.LoadConfig("../../.secret/config.json")
	if err != nil {
		log.Fatal(err)
	}
	println(config.Crowemi.ClientName)
}
