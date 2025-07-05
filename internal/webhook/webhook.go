package webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Luiso9/goclaim-soccer/internal/api"

	discordwebhook "github.com/bensch777/discord-webhook-golang"
)


func Notify(result *api.ClaimResponse) {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	card := result.Card

	embed := discordwebhook.Embed{
		Title:     "Claim Success",
		Color:     5763719,
		Timestamp: time.Now(),
		Thumbnail: discordwebhook.Thumbnail{
			Url: "https://mirrorcdn.soccerguru.live/cards/master/" + card.Uuid + ".png",
		},
		Fields: []discordwebhook.Field{
			discordwebhook.Field{
				Name:   "Name",
				Value:  card.FullName,
				Inline: false,
			},
			discordwebhook.Field{
				Name:   "Position",
				Value:  card.Position,
				Inline: false,
			},
			discordwebhook.Field{
				Name:   "Sell Value",
				Value:  fmt.Sprintf("%d", card.Bin),
				Inline: false,
			},
		},
	}
	err := SendNotify(webhookUrl, embed)
	if err != nil {
		log.Println("Discord webhook send failed", err)
	}
}

func NotifyDaily(dailyRespone *api.DailyResponse) {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	embed := discordwebhook.Embed{
		Title:     "Daily Login Success",
		Color:     5763719,
		Timestamp: time.Now(),
		Thumbnail: discordwebhook.Thumbnail{
			Url: "https://cdn.driannsa.my.id/megalon2d-PeakCinema-0.1.0.png.256x256_q95_crop.png",
		},
		Fields: []discordwebhook.Field{
			{
				Name:   "Streak",
				Value:  fmt.Sprintf("%d", dailyRespone.Streak),
				Inline: true,
			},
			{
				Name:   "Pack Reward",
				Value:  fmt.Sprintf("%t", dailyRespone.IsPackReward),
				Inline: true,
			},
		},
	}
	err := SendNotify(webhookUrl, embed)
	if err != nil {
		log.Println("Discord webhook send failed", err)
	}
}

func SendNotify(link string, embeds discordwebhook.Embed) error {
	hook := discordwebhook.Hook{
		Username: "GoClaim",
		Embeds:   []discordwebhook.Embed{embeds},
	}

	payload, err := json.Marshal(hook)
	if err != nil {
		log.Fatal(err)
	}

	return discordwebhook.ExecuteWebhook(link, payload)

}
