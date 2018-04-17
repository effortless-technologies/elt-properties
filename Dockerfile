#FROM golang:latest
#
#ENV MONGO_ADDR='104.198.34.190:27017'
#
#ADD . /go/src/app
#WORKDIR /go/src
#RUN go get app
#RUN go install app
#
#ENTRYPOINT /go/bin/app -mongoAddr=$MONGO_ADDR

FROM iron/go:dev

# ENV MONGO_ADDR='104.198.34.190:27017'
ENV MONGO_ADDR='prod'

WORKDIR /app

# Build API
ENV SRC_DIR=/go/src/github.com/effortless-technologies/elt-properties
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get
RUN cd $SRC_DIR; go build -o api; cp api /app/

ENTRYPOINT ./api -mongoAddr=$MONGO_ADDR
