package notify

import (
	"log"
	"os"
	"time"

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
	} `yaml:"Distinguished Club Program Report"`
}

func LoadDcpReport() {
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
}

func GetSummary() string {
	summary := "Achieved " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target
	if dcpReport.DCPReport.DCPStatus.Overall.Distinguished == "Yes" {
		summary = "Achieved Distinguished Status " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target + " goals"
	}
	if dcpReport.DCPReport.DCPStatus.Overall.SelectDistinguished == "Yes" {
		summary = "Achieved Select Distinguished Status " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target + " goals"
	}
	if dcpReport.DCPReport.DCPStatus.Overall.PresidentsDistinguished == "Yes" {
		summary = "Achieved President's Distinguished Status " + dcpReport.DCPReport.DCPStatus.Overall.Current + " of " + dcpReport.DCPReport.DCPStatus.Overall.Target + " goals"
	}
	return summary
}

func GetMessage() string {
	// Open the YAML file
	yaml, err := os.ReadFile("../reports/dcp_report.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	currentTime := time.Now()

	question := util.QuestionBank[int(currentTime.Month())] + string(yaml)
	return util.Chat(question + os.Getenv("CHAT_OUTPUT_FORMAT_PROMPT"))
}
