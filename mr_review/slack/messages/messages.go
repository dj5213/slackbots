package messages

import (
	"log"

	"github.com/nlopes/slack"
)

type IncomingSlackMessage struct {
	Channel string
	Text    string
	TS      string
	User    string
}

func SendMessage(params slack.PostMessageParameters, msg string,
	channel string, api *slack.Client) {
	log.Printf("Posting to channel: %s\n", channel)

	channelID, timestamp, err := api.PostMessage(channel, msg, params)
	if err != nil {
		log.Printf("[ERROR]: issue sending Slack message: %s", err)
	} else {
		log.Printf("Message posted to channel %s at %s", channelID, timestamp)
	}
}
