package notify

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

var CorporateTerms = []map[string]string{
	{"tm_term": "Distinguished Club Program", "corp_term": "Strategic plan"},
	{"tm_term": "Membership", "corp_term": "Particiption"},
	{"tm_term": "Membership-renewal", "corp_term": "Renewal"},
	{"tm_term": "Pathways", "corp_term": "Curriculum"},
	{"tm_term": "Officer", "corp_term": "Executive Board"},
	{"tm_term": "Dues", "corp_term": "Tution"},
	{"tm_term": "Goal", "corp_term": "Critical Success Factor"},
	{"tm_term": "DTM", "corp_term": "Milestone"},
	{"tm_term": "Member", "corp_term": "Participant"},
	{"tm_term": "Club", "corp_term": "Program"}}

func LoadDcpReport() {
	// Open the YAML file
	decodeDcpReport("../reports/dcp_report.yaml")
}

func decodeDcpReport(dcp_report_yaml_path string) {
	// Open the YAML file
	file, err := os.Open(dcp_report_yaml_path)
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

func ReplaceWithCorpTerms(message_body string) string {
	corp_message_body := message_body
	for idx, corp_value := range CorporateTerms {
		fmt.Printf("Corp Term #%d: Replacing %s with %s\n", idx+1, corp_value["tm_term"], corp_value["corp_term"])
		corp_message_body = caseInsensitiveReplacer(corp_message_body, corp_value["tm_term"], corp_value["corp_term"])
	}
	return corp_message_body
}

func caseInsensitiveReplacer(message string, toReplace string, replaceWith string) string {
	pattern := fmt.Sprintf(
		`(?i)(^|[^-])(%s)($|[^-])`,
		toReplace,
	)
	var regx = regexp.MustCompile(pattern)
	escapedReplaceWith := regexp.QuoteMeta(replaceWith)
	return regx.ReplaceAllString(message, `${1}`+escapedReplaceWith+`${3}`)
}
