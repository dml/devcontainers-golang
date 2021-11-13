FROM golang:1.17-alpine

RUN apk add --update --no-cache \
  sudo openssh-client tzdata curl git

WORKDIR /go/src/app

RUN go install golang.org/x/tools/gopls@latest
