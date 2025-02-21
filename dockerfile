FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/api/*.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
