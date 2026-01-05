package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	notify "github.com/d2tm/tmi-status-checker/util/notify"
)

func TestGenerateMessage(t *testing.T) {
	notify.LoadDcpReport()
	message_card := GenerateMessageCard()
	payload, err := json.MarshalIndent(message_card, "", "	")
	if err != nil {
		log.Fatalf("Failed to get JSON payload: %v", err)
	}

	fmt.Printf("Message Body: %s", bytes.NewBuffer(payload))
}
