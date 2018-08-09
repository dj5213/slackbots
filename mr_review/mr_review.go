// +build linux

package main

import (
	"flag"
	"log"
	"os"

	"github.com/nlopes/slack"

	"github.com/dj5213/slackbots/mr_review/bot"
	"github.com/dj5213/slackbots/mr_review/gitlab"
	"github.com/dj5213/slackbots/mr_review/slack/messages"
)

func main() {
	channel := flag.String("reply_channel", "", "Channel for sending messages to")
	baseURL := flag.String("custom_gitlab_url", "", "Custom gitlab URL (if not gitlab.com)")
	flag.Parse()

	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatalf("No Slack token found - add token to env variable SLACK_TOKEN")
	}

	api := slack.New(token)

	_, err := api.AuthTest()
	if err != nil {
		log.Fatal(err)
	}

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	itemchan := make(chan *messages.IncomingSlackMessage)
	gitlabClient := gitlab.NewGitlabClient(*baseURL)

	go bot.ProcessForever(itemchan, api, gitlabClient)
	if *channel != "" {
		go bot.ReadInputForever(api, *channel)
	}
	log.Printf("Mr. Review loaded up")

	for {
		select {
		case item := <-rtm.IncomingEvents:
			switch ev := item.Data.(type) {
			case *slack.MessageEvent:
				// Only care about messages not made by the bot
				if ev.User != bot.Username {
					msg := messages.IncomingSlackMessage{
						Text:    ev.Text,
						Channel: ev.Channel,
						TS:      ev.Timestamp,
						User:    ev.User,
					}
					itemchan <- &msg
				}
			}
		}
	}
}
