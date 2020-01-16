FROM golang:1.13.4-alpine3.10 as builder

WORKDIR /go/src/app

COPY ./src/app/ .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine

COPY --from=builder /go/bin/app /go/bin/app

CMD [ "/go/bin/app" ]