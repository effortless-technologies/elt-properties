FROM golang:latest

ENV MONGO_ADDR='104.198.34.190:27017'

ADD . /go/src/app
WORKDIR /go/src
RUN go get app
RUN go install app

ENTRYPOINT /go/bin/app -mongoAddr=$MONGO_ADDR
