FROM golang:1.23.7-bookworm

WORKDIR /chitchat

COPY go.mod go.sum  ./
RUN go mod download


COPY . .
RUN go build -v -o /usr/local/bin/chitchat ./cmd/server/main.go

CMD ["chitchat"]
