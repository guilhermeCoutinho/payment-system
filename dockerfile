FROM golang:1.15 as builder

ADD . /go/src/github.com/guilhermeCoutinho/payment-system
WORKDIR /go/src/github.com/guilhermeCoutinho/payment-system
RUN mkdir bin/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=readonly -o bin/payment-system

FROM alpine:3.8

COPY --from=builder /go/src/github.com/guilhermeCoutinho/payment-system/bin/payment-system /app/payment-system
COPY --from=builder /go/src/github.com/guilhermeCoutinho/payment-system/config /app/config

WORKDIR /app
CMD /app/payment-system serve
