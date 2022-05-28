FROM golang:1.18.1-stretch

WORKDIR /app
ADD ./main.go /app
CMD go run main.go
