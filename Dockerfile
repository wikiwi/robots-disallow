FROM golang:1.6-alpine

RUN apk add --no-cache \
    git \
    openssh-client

WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download && \
    go-wrapper install

ENTRYPOINT ["app"]
