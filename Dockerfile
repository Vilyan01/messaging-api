FROM golang

ADD . /go/src/github.com/Vilyan01/messaging-api

RUN go install github.com/Vilyan01/messaging-api

ENTRYPOINT /go/bin/messaging-api

EXPOSE 8080
