FROM golang:1.17-alpine

RUN apk add --update --no-cache \
  build-base \
  sudo openssh-client tzdata curl git

WORKDIR /go/src/app

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/go-delve/delve/cmd/dlv@master
RUN go install honnef.co/go/tools/cmd/staticcheck@latest