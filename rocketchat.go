package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
  "os"
)


func SendRocketChatAlert() {

  rocketChatURL := os.Getenv("ROCKET_CHAT_URL")
	msg := "No new deposits within one hour after the latest invoice!"
	payload := []byte(fmt.Sprintf(`{"text": "%s"}`, msg))

	// Make a POST request to Rocket.Chat webhook URL
	resp, err := http.Post(rocketChatURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Error sending Rocket.Chat alert:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to send Rocket.Chat alert. Status Code:", resp.StatusCode)
		return
	}
	log.Println("Rocket.Chat alert sent successfully!")
}

