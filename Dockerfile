FROM golang:latest
WORKDIR /src
COPY . .
CMD go run main.go