FROM golang

RUN go get github.com/conslo/badlink
WORKDIR $GOPATH/src/github.com/conslo/badlink
RUN go install -a
RUN go test
