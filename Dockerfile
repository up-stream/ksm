FROM golang

WORKDIR /go/src/github.com/up-stream/ksm

ADD . /go/src/github.com/up-stream/ksm

RUN go get -t -v ./...

CMD ["go", "run", "example/basic.go"]

EXPOSE 8080