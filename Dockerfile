FROM alpine:3.1
MAINTAINER Eric Holmes <eric@remind101.com>

COPY ./ /go/src/github.com/remind101/acme-inc
COPY ./bin/build /build
RUN /build
WORKDIR /go/src/github.com/remind101/acme-inc

CMD ["acme-inc", "server"]
