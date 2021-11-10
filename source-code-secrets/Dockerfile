FROM golang:1.15.1-alpine3.12 AS build-env

WORKDIR /tmp/simple-go-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine:3.13

EXPOSE 8080

RUN apk add --no-cache ca-certificates bash

COPY --from=build-env /tmp/simple-go-app/gitops-secrets-sample-app /app/gitops-secrets-sample-app

COPY settings.ini /config/settings.ini

WORKDIR /app


CMD ["./gitops-secrets-sample-app"]
