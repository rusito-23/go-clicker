FROM golang:1.16-alpine

# Setup files
WORKDIR /app
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

# Expose port
EXPOSE 8080

# Entry point specification
ENTRYPOINT CompileDaemon \
--build="go build -o /go-friends" \
--command=/go-friends
