package scheduler

import (
	"errors"
	"io"
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
	// Hourly Claim
	for _, req := range api.OptionRequest() {
		result, err := api.SendRequest(req)

		// Skipping response from OPTIONS request.
		if req.Method == "OPTIONS" {
			continue
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Claimed")
			} else {
				log.Println("API request failed:", err)
			}
			continue
		}

		webhook.Notify(result)
		log.Printf("Claimed: %s (%s) BIN: %d", result.Card.FullName, result.Card.Position, result.Card.Bin)
	}

	// Daily Claim
	for _, req := range api.DailyLoginRequest() {
		dailyRespone, err := api.DailyRequest(req)

		// ignoring options request
		if req.Method == "OPTIONS" {
			continue
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Claimed")
			} else {
				log.Println("API request for daily failed:", err)
			}
			continue
		}

		if dailyRespone != nil {
			webhook.NotifyDaily(dailyRespone)
			log.Printf("Daily login streak: %d, Pack reward: %t", dailyRespone.Streak, dailyRespone.IsPackReward)
		} else {
			log.Println("Daily response is nil, skipping NotifyDaily")
		}
	}
}
