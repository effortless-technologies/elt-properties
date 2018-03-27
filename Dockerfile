FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/effortless-technologies/elt-properties
ADD . $SRC_DIR
RUN cd #SRC_DIR; go get
RUN cd $SRC_DIR; go build -o api; cp api /app/

ENTRYPOINT ["./api"]