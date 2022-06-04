FROM golang:1.18.1-stretch

WORKDIR /app
ADD . .
RUN go mod download
EXPOSE 8080

CMD go run main.go
