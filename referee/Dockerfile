FROM golang:latest

COPY . /go/src/github.com/BarryMcAuley/golden_goal/referee
WORKDIR /go/src/github.com/BarryMcAuley/golden_goal/referee

RUN go get -v
RUN go build cmd/referee_server.go
RUN cp referee_server /go/bin/

CMD ["/go/bin/referee_server", "-dbhost", "rethink"]
