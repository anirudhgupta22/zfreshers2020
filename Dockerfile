FROM golang:1.12.1-alpine as builder

# setup the working directory
WORKDIR /go/src/zfreshers2020

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello-app-linux-amd64

ENTRYPOINT ["./hello-app-linux-amd64"]
