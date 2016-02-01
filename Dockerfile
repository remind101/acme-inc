FROM golang:1.5.3-alpine
MAINTAINER Eric Holmes <eric@remind101.com>

COPY ./ /go/src/github.com/remind101/acme-inc
RUN go install github.com/remind101/acme-inc
WORKDIR /go/src/github.com/remind101/acme-inc

CMD ["acme-inc", "server"]
