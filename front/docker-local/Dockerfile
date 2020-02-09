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

# vue.js用
## FIXME: webpackでvueをコンパイルできるようにする
## FIXME: CDNで代用するのでpackage.jsonは休止
#RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
#RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
#RUN apt-get -y update && apt-get -y install yarn
#RUN yarn install

COPY front/static /static
COPY front/template /template

ENTRYPOINT ["/go/bin/front"]
