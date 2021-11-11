FROM golang:1.16-alpine3.13 AS build-env

WORKDIR /tmp/workdir

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine:3.13

EXPOSE 8080

RUN apk add --no-cache ca-certificates bash

COPY --from=build-env /tmp/workdir/static /app/static
COPY --from=build-env /tmp/workdir/canary-app /app/canary-app

WORKDIR /app

CMD ["./canary-app"]
