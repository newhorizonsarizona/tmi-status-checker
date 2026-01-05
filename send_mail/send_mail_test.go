package main

import (
	"fmt"
	"os"
	"testing"

	notify "github.com/d2tm/tmi-status-checker/util/notify"
)

func TestGenerateMessage(t *testing.T) {
	notify.LoadDcpReport()
	club_number := os.Getenv("CLUB_NUMBER")
	club_name := os.Getenv("CLUB_NAME")
	email_body := generateMessageBody(club_number, club_name)
	fmt.Printf("Email Body: %s", email_body)
}
