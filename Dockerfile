# Start from the official Go image
FROM golang:1.23.3-alpine
RUN go install github.com/air-verse/air@latest
WORKDIR /app
ENTRYPOINT ["air"]