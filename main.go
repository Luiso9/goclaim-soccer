package main

import (
	"log"
	
	// "github.com/Luiso9/goclaim-soccer/internal/api"
	// "github.com/Luiso9/goclaim-soccer/internal/webhook"
	"github.com/Luiso9/goclaim-soccer/internal/scheduler"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load() // defaults to .env in root
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Commented code used to debug the discord webhook :D

	// mock := &api.ClaimResponse{
	// 	Card: struct {
	// 		FullName string `json:"full_name"`
	// 		Position string `json:"position"`
	// 		Rating   int    `json:"rating"`
	// 		Value    int    `json:"value"`
	// 		Bin      int    `json:"bin"`
	// 	}{
	// 		FullName: "Test Player",
	// 		Position: "CM",
	// 		Rating:   99,
	// 		Value:    999999,
	// 		Bin:      666666,
	// 	},
	// }

	// webhook.Notify(mock)
	scheduler.DoHourlyJob()
}
