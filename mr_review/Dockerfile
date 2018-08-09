FROM golang

COPY mr_review /go/bin/

COPY bot/templates/* /go/src/mr_review/templates/slack_msg/

COPY config/* /go/src/mr_review/config/

ARG slack_token
ENV SLACK_TOKEN $slack_token

ARG gitlab_token
ENV GITLAB_TOKEN $gitlab_token

ENTRYPOINT /go/bin/mr_review 2> output.log
