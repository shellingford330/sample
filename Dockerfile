FROM golang:1.16

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /bin/main

CMD ["/bin/main"]
