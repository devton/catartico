package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/pg.v3"
)

// MessageTemplate templating for sent messages to slack
const (
	MessageTemplate = `{
    "channel": "%s",
    "username": "catartico",
    "icon_url": "https://s3.amazonaws.com/cdn.catarse/assets/catartico72.png",
    "attachments": [
    		{
    				"title": "%s",
    				"text": "%s",
    				"mrkdwn_in": ["text", "title"]
    		}
    ]}`
)

// EventMessage type is the event struct received though listen
type EventMessage struct {
	Channel string
	Name    string
	Title   string
	Text    string
}

func main() {
	db := pg.Connect(&pg.Options{
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_DATABASE"),
	})

	defer db.Close()

	ln, err := db.Listen(os.Getenv("LISTEN_BUCKET"))
	if err != nil {
		panic(err)
	}

	wait := make(chan string)

	// repceiving pg_notify
	for {
		go func() {
			_, payload, _ := ln.Receive()
			wait <- payload
		}()

		select {
		case payload := <-wait:
			// post payload to slack
			go func() {
				var event = &EventMessage{}
				json.Unmarshal([]byte(payload), &event)

				if event.Name != "" {
					botMessage := []byte(
						fmt.Sprintf(
							MessageTemplate,
							event.Channel, event.Title, event.Text))

					fmt.Printf(string(botMessage))
					http.Post(
						os.Getenv("SLACK_HOOK_URL"),
						"application/json", bytes.NewReader(botMessage))
				}
			}()

		}
		time.Sleep(50 * time.Millisecond)
	}

}
