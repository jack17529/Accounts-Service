FROM golang

RUN mkdir -p /go/src/github.com/faith/Accounts2

ADD . /go/src/github.com/faith/Accounts2

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/faith/Accounts2/accounts/cmd  -watch github.com/faith/Accounts2/accounts
