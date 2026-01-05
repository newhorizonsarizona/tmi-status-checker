package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	notify "github.com/d2tm/tmi-status-checker/util/notify"
)

func main() {
	notify.LoadDcpReport()

	//Send the announcement
	err := sendAnnouncement(GenerateMessageCard())
	if err != nil {
		log.Fatalf("Failed to send DCP Report announcement message: %v", err)
	} else {
		log.Println("DCP Report Announcement message was sent successfully!")
	}

}

func GenerateMessageCard() map[string]interface{} {
	statusSubheading := notify.GetSummary()
	statusMessage := "Congratulations on your achievements!"
	statusMessage = notify.GetMessage()

	clubNumber := os.Getenv("CLUB_NUMBER")
	messageCard := map[string]interface{}{
		"type": "message",
		"attachments": []map[string]interface{}{
			{
				"contentType": "application/vnd.microsoft.card.adaptive",
				"content": map[string]interface{}{
					"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
					"type":    "AdaptiveCard",
					"version": "1.3",
					"body": []map[string]interface{}{
						{
							"type":   "TextBlock",
							"text":   "Club DCP Status",
							"weight": "bolder",
							"size":   "extraLarge",
							"color":  "accent",
							"wrap":   true,
						},
						{
							"type":    "TextBlock",
							"text":    statusSubheading,
							"weight":  "bolder",
							"size":    "large",
							"spacing": "medium",
							"wrap":    true,
						},
						{
							"type": "TextBlock",
							"text": statusMessage,
							"wrap": true,
						},
						{
							"type":        "Image",
							"url":         "https://raw.githubusercontent.com/newhorizonsarizona/tmi-status-checker/refs/heads/main/reports/dcp_report.png",
							"pixelWidth":  400,
							"pixelHeight": 500,
						},
					},
					"actions": []map[string]interface{}{
						{
							"type":  "Action.OpenUrl",
							"title": "Learn More",
							"url":   "https://dashboards.toastmasters.org/ClubReport.aspx?id=" + clubNumber,
						},
					},
				},
			},
		},
	}

	return messageCard
}

func sendAnnouncement(cardContent map[string]interface{}) error {
	channelWebhookUrl := os.Getenv("CHANNEL_WEBHOOK_URL")
	// Create the message payload
	payload, err := json.Marshal(cardContent)
	if err != nil {
		return err
	}

	// Send the POST request
	resp, err := http.Post(channelWebhookUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil

}
