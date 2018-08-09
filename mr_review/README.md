# Mr Review Slack bot

This is a Slack bot that hooks up to specified Gitlab repos and tells you info about open merge requests. Right now it just says how many there are open for each given project, but there are more options/commands on the way.

# Running Mr Review:

- export `SLACK_TOKEN` to env
- export `GITLAB_TOKEN` to env
- Add the bot username to `constants.go`
- Add project name and ID into `projects.yaml`
- `go run mr_review.go`

# Running in Docker container:

Follow the steps above and (except exporting the tokens) and:
- Build go binary: `go build mr_review.go`
- Run `docker build` and pass in your Slack + Gitlab tokens
    - Example:
    `docker build --build-arg gitlab_token='<GITLAB_TOKEN>' --build-arg slack_token='<SLACK_TOKEN' -t mr_review .`
- `docker run mr_review:latest`
