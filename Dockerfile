FROM remind101/go:1.4
MAINTAINER Eric Holmes <eric@remind101.com>

WORKDIR /go/src/github.com/remind101/acme-inc
COPY ./ /go/src/github.com/remind101/acme-inc
RUN go install ./...
ENTRYPOINT ["/go/bin/acme-inc"]

EXPOSE 8080
