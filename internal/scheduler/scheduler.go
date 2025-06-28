package scheduler

import (
	"log"
	"time"

	"github.com/Luiso9/goclaim-soccer/internal/api"
	"github.com/Luiso9/goclaim-soccer/internal/webhook"
)

func DoHourlyJob() {
	go run()
	ticker := time.NewTicker(1 * time.Hour)
	for range ticker.C {
		run()
	}
}

func run() {
	for _, req := range api.OptionRequest() {
		result, err := api.SendRequest(req)

		if err != nil {
			log.Println("API request failed :", err)
			continue
		}

		// Skipping response from OPTIONS request.
		if req.Method == "OPTIONS" {
			log.Println("Empty")
			continue
		}

		webhook.SendNotify(result)
		log.Printf("Claimed: %s (%s) BIN: %d", result.Card.FullName, result.Card.Position, result.Card.Bin)
	}
}
