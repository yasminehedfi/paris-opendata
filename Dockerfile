FROM golang:1.25 as builder

WORKDIR /app
COPY backend/ ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["./main"]
