package bot

import (
	"bytes"
	"log"
	"os"

	"text/template"

	botGitlab "github.com/dj5213/slackbots/mr_review/gitlab"
)

// Dockerized path to Mr Review templates
var slackMsgPath string = "/src/mr_review/templates/slack_msg/"

func loadTemplate(path string, data interface{}) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working dir: %s", err)
		return "", err
	}

	tmpl, err := template.ParseFiles(dir + path)
	if err != nil {
		log.Printf("Error rendering template: %s", err)
		return "", err
	}

	var b bytes.Buffer

	if err := tmpl.Execute(&b, data); err != nil {
		log.Printf("Issue executing %s template: %s", path, err)
	}
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func formatMRs(mrs botGitlab.ActiveProjects) (string, error) {
	return loadTemplate(slackMsgPath+"open_mrs.tmpl", mrs)
}

func formatAvailCommands() (string, error) {
	return loadTemplate(slackMsgPath+"commands.tmpl", getCmds())
}
