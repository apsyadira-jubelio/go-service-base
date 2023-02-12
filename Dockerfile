# Builder
FROM golang:1.19.4-alpine3.17 as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /apps
ENV CGO_ENABLED=0

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /apps

WORKDIR /apps

EXPOSE 8083

COPY --from=builder /apps/chat-webhook /apps/

WORKDIR /apps


CMD /apps/chat-webhook
