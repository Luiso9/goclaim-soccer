package main

import (
	// "github.com/Luiso9/goclaim-soccer/internal/api"
	"github.com/Luiso9/goclaim-soccer/internal/scheduler"
	// "github.com/Luiso9/goclaim-soccer/internal/webhook"
)

func main() {
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
