FROM golang

ADD . /go

WORKDIR /go

ENV GOPATH /go

ENV GOBIN /go/bin

RUN ["go","install","src/main.go"]

ENTRYPOINT ["/go/bin/main"]

EXPOSE 8080  
