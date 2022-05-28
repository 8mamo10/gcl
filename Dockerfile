FROM golang:1.18.1-stretch

WORKDIR /go/src
ADD ./main.go /go/src
CMD go run main.go
