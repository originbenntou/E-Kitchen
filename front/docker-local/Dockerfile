FROM golang:latest

ENV WORK_DIR /go/src/github.com/originbenntou/E-Kitchen
ENV MAIN_PATH $WORK_DIR/front

WORKDIR $WORK_DIR

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $MAIN_PATH

COPY front/template /template

ENTRYPOINT ["/go/bin/front"]
