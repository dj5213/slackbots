package bot

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	botGitlab "github.com/dj5213/slackbots/mr_review/gitlab"
	"github.com/dj5213/slackbots/mr_review/slack/messages"

	"github.com/nlopes/slack"
	"github.com/xanzy/go-gitlab"
)

func replyToMSg(msg string, channel string,
	api *slack.Client, client *gitlab.Client) {
	params := slack.PostMessageParameters{
		Username: "Mr. Review",
		AsUser:   true,
	}
	switch {

	case strings.Contains(msg, "open mr"), strings.Contains(msg, "open merge request"):
		projs := botGitlab.GetOpenMRs(client)

		mrs, err := formatMRs(projs)
		if err != nil {
			log.Printf("Error formatting MR template: %s", err)
		} else {
			messages.SendMessage(params, mrs, channel, api)
		}

	default:
		cmds, err := formatAvailCommands()

		if err != nil {
			log.Printf("Error formatting available commands: %s", err)
		}

		if cmds == "" {
			log.Printf("Not sending Slack message, empty text")
		} else {
			messages.SendMessage(params, cmds, channel, api)
		}
	}
}

func ProcessForever(incoming chan *messages.IncomingSlackMessage, api *slack.Client, client *gitlab.Client) {
	ticker := time.Tick(250 * time.Millisecond)
	msgs := make([]*messages.IncomingSlackMessage, 0, 10)
	for {
		select {
		case <-ticker:
			newMsgs := make([]*messages.IncomingSlackMessage, 0, len(msgs))
			for i := range msgs {
				ProcessMessage(msgs[i], api, client)
			}
			msgs = newMsgs
		case msg := <-incoming:
			msgs = append(msgs, msg)
		}
	}
}

func ReadInputForever(api *slack.Client, channel string) {
	params := slack.PostMessageParameters{
		Username: "Mr. Review",
		AsUser:   true,
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		messages.SendMessage(params, text, channel, api)
	}
}

func ProcessMessage(msg *messages.IncomingSlackMessage, api *slack.Client,
	client *gitlab.Client) {
	log.Printf("Message received[%s]: %s", msg.Channel, msg.Text)
	lowerMsg := strings.ToLower(msg.Text)

	if strings.Contains(msg.Text, Username) {
		// @mentioned the bot
		replyToMSg(lowerMsg, msg.Channel, api, client)
	} else if string(msg.Channel[0]) == "D" {
		// Direct message
		replyToMSg(lowerMsg, msg.Channel, api, client)
	} else {
		log.Printf("Message received, but not for me")
	}
}
