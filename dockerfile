FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o main ./cmd/api/*.go

CMD ["./main"]
