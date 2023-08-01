FROM golang:1.20.6-alpine3.18

WORKDIR /usr/src/app

RUN apk add build-base

ENV CGO_ENABLED 1

RUN export GO111MODULE=on

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -tags musl -o build/ cmd/main.go

CMD ["build/main"]