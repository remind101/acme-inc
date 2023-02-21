FROM golang:1.19-alpine

WORKDIR /usr/src/acme-inc

CMD ["acme-inc", "server"]

COPY . .
RUN go build -v -o /usr/local/bin/acme-inc ./...
