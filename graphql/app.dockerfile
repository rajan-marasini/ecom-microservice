FROM golang:1.25-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/rajan-marasini/ecom-microservice
COPY go.mod go.sum ./
COPY vendor vendor
COPY catalog catalog
COPY account account
COPY order order
COPY graphql graphql
RUN GO111MODULE=on CGO_ENABLED=0 \
    go build -mod=vendor -o app ./account/cmd/account


FROM alpine:latest

WORKDIR /usr/bin

COPY --from=build /go/bin .

EXPOSE 8080

CMD [ "app" ]
