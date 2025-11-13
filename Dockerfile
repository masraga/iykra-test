FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod download
# RUN go install github.com/air-verse/air@latest

RUN go build -o tmp/main cmd/main.go

CMD ["tmp/main"]