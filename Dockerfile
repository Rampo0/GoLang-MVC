FROM golang:latest

WORKDIR /go/src/go_framework

COPY . .

RUN go get -d -v ./...

# RUN go build -o server main.go

# RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN go install -v ./...

CMD ["go_framework"]
