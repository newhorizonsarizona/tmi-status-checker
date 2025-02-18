package main

import (
	"fmt"
	"os"

	notify "github.com/d2tm/tmi-status-checker/util/notify"
	"gopkg.in/gomail.v2"
)

func main() {
	notify.LoadDcpReport()
	m := gomail.NewMessage()
	email_from := os.Getenv("EMAIL_ID")
	email_password := os.Getenv("EMAIL_PASSWORD")
	club_number := os.Getenv("CLUB_NUMBER")
	club_name := os.Getenv("CLUB_NAME")
	//email_to := "members-" + club_number + "@toastmastersclubs.org"
	email_to := "anand_vijayan@yahoo.com" //TODO: Remove after testing
	email_subject := notify.GetSummary()
	email_body := generateMessageBody(club_number, club_name)
	m.SetHeader("From", email_from)
	m.SetHeader("To", email_to)
	m.SetHeader("Subject", "DCP Report: "+email_subject)
	m.SetBody("text/plain", email_body)
	m.Attach("../reports/dcp_report.png")
	d := gomail.NewDialer("smtp.gmail.com", 587, email_from, email_password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func generateMessageBody(club_number string, club_name string) string {
	if club_name == "" {
		club_name = "Toastmasters Club"
	}
	return fmt.Sprintf(`
Dear %s Members,

%s

https://dashboards.toastmasters.org/ClubReport.aspx?id=%s

Best Regards,
Executive Team
	`, club_name, notify.GetMessage(), club_number)
}
