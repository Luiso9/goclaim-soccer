package webhook

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Luiso9/goclaim-soccer/internal/api"

	discordwebhook "github.com/bensch777/discord-webhook-golang"
)

var webhookUrl = "https://discord.com/api/webhooks/1283840202313306144/Hvr5sbMa7TK_lvmIxdXY5QEhn37_ET1CQ5uU2s7e6Am0iWhLNnwdkix-HeMVlLpZ0ee7"

func Notify(result *api.ClaimResponse) {
	card := result.Card

	embed := discordwebhook.Embed{
		Title:     "Claim Success",
		Color:     5763719,
		Timestamp: time.Now(),
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
				Name:   "Bin",
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
