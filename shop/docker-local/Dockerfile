FROM golang:latest as builder

ENV WORK_DIR /go/src/github.com/originbenntou/E-Kitchen
ENV MAIN_PATH $WORK_DIR/shop

WORKDIR $WORK_DIR

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $MAIN_PATH

ENTRYPOINT ["/go/bin/shop"]
