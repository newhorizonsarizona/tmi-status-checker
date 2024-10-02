package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	util "github.com/newhorizonsarizona/tmi-status-checker/util"

	"gopkg.in/yaml.v3"
)

// Create an instance of the Config struct
var dcpReport DCPReport

// Define a struct that matches the structure of your YAML file
type DCPReport struct {
	DCPReport struct {
		Administration struct {
			ClubOfficerListOnTime struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Club officer list on time"`
			MembershipRenewalDuesOnTime struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Membership-renewal dues on time"`
		} `yaml:"Administration"`
		DCPStatus struct {
			Membership struct {
				Base     string `yaml:"Base"`
				Required string `yaml:"Required"`
				ToDate   string `yaml:"To Date"`
			} `yaml:"Membership"`
			Overall struct {
				Year                    string `yaml:"Year"`
				Current                 string `yaml:"Current"`
				Distinguished           string `yaml:"Distinguished"`
				SelectDistinguished     string `yaml:"Select Distinguished"`
				PresidentsDistinguished string `yaml:"President's Distinguished"`
				Target                  string `yaml:"Target"`
			} `yaml:"Overall"`
		} `yaml:"DCP Status"`
		Education struct {
			Level1Awards struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Level 1 awards"`
			Level2Awards struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Level 2 awards"`
			MoreLevel2Awards struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"More Level 2 awards"`
			Level3Awards struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Level 3 awards"`
			Level4Level5OrDTMAward struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Level 4; Level 5; or DTM award"`
			OneMoreLevel4Level5OrDTMAward struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"One more Level 4; Level 5; or DTM award"`
		} `yaml:"Education"`
		Membership struct {
			NewMembers struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"New members"`
			MoreNewMembers struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"More new members"`
		} `yaml:"Membership"`
		Training struct {
			ClubOfficersTrainedJunToAug struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Club officers trained June-August"`
			ClubOfficersTrainedNovToFeb struct {
				Achieved string `yaml:"achieved"`
				Status   string `yaml:"status"`
				Target   string `yaml:"target"`
			} `yaml:"Club officers trained November-February"`
		} `yaml:"Training"`
	} `yaml:"DCP Report"`
}

func main() {
	// Open the YAML file
	file, err := os.Open("../reports/dcp_report.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	// Create a new decoder
	decoder := yaml.NewDecoder(file)

	// Decode the YAML into the struct
	err = decoder.Decode(&dcpReport)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Access the values from the config struct
	err = sendAnnouncement(generateAdaptiveCard())
	if err != nil {
		log.Fatalf("Failed to send announcement: %v", err)
	} else {
		log.Println("Announcement was sent successfully!")
	}
}

func generateAdaptiveCard() map[string]interface{} {
	statusSubheading := "Achieved " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target
	statusMessage := "Congratulations on your achievements!"
	if dcpReport.DCPReport.DCPStatus.Overall.Distinguished == "Yes" {
		statusSubheading += "Achieved Distinguished Status " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target + " goals"
	}
	if dcpReport.DCPReport.DCPStatus.Overall.SelectDistinguished == "Yes" {
		statusSubheading += "Achieved Select Distinguished Status " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target + " goals"
	}
	if dcpReport.DCPReport.DCPStatus.Overall.PresidentsDistinguished == "Yes" {
		statusSubheading += "Achieved President's Distinguished Status " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target + " goals"
	}

	// Open the YAML file
	yaml, err := os.ReadFile("../reports/dcp_report.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	question := `
				In one concise to the point paragraph with a formal encouraging tone highlight the club achievements. 
				The first part praises the club for their overall achievements and membership, and 
				the second part commends the club for DCP goals achieved in each category.
				` + string(yaml)
	statusMessage = util.Chat(question)

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
							"type":        "Image",
							"url":         "https://github.com/newhorizonsarizona/tmi-status-checker/blob/main/dcp_report.png?raw=true",
							"pixelWidth":  400,
							"pixelHeight": 500,
						},
						{
							"type": "TextBlock",
							"text": statusMessage,
						},
					},
					"actions": []map[string]interface{}{
						{
							"type":  "Action.OpenUrl",
							"title": "Learn More",
							"url":   "https://dashboards.toastmasters.org/ClubReport.aspx?id=00006350",
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
