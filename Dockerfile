FROM golang:1.18.1-stretch

WORKDIR /workspace
ADD go.mod /workspace
RUN go mod download

CMD go run main.go
