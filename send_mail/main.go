package main

import (
	"fmt"
	"os"
	"strings"

	notify "github.com/d2tm/tmi-status-checker/util/notify"
	"gopkg.in/gomail.v2"
)

func main() {
	notify.LoadDcpReport()
	m := gomail.NewMessage()
	email_user := os.Getenv("EMAIL_USER")
	email_password := os.Getenv("EMAIL_PASSWORD")
	club_number := os.Getenv("CLUB_NUMBER")
	club_name := os.Getenv("CLUB_NAME")
	email_from := os.Getenv("EMAIL_FROM")
	email_to := os.Getenv("EMAIL_TO")
	email_cc := os.Getenv("EMAIL_CC")
	//email_to := "anand_vijayan@yahoo.com" //TODO: Remove after testing
	email_subject := notify.GetSummary()
	email_body := generateMessageBody(club_number, club_name)
	m.SetHeader("From", email_from)
	m.SetHeader("To", email_to)
	cc_addrs := strings.Split(email_cc, ",")
	if len(cc_addrs) > 0 {
		m.SetHeader("Cc", cc_addrs...)
	}
	m.SetHeader("Subject", "DCP Report ("+club_name+"): "+email_subject)
	m.SetBody("text/html", email_body)
	m.Attach("../reports/dcp_report.png")
	d := gomail.NewDialer("email-smtp.us-west-2.amazonaws.com", 587, email_user, email_password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func generateMessageBody(club_number string, club_name string) string {
	if club_name == "" {
		club_name = "Toastmasters Club"
	}
	message := notify.GetMessage()
	message = strings.ReplaceAll(message, "```html", "")
	message = strings.ReplaceAll(message, "```", "")
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Club Achievements Overview</title>
</head>
<body>
Dear <b>%s</b> Members,<br/>
<br/>
Thank you for everything you are doing for the club. The toastmaster year runs from 1st July through 30th June
and it behooves all members to help the club succeed. Here is this month's distinguished club program (DCP) report.
<br/>
%s
View the club <a href='https://dashboards.toastmasters.org/ClubReport.aspx?id=%s'>DCP Report</a>.<br/>
<br/>
Best Regards,<br/>
<b>Executive Team</b>
</body>
</html>
`, club_name, message, club_number)
}
