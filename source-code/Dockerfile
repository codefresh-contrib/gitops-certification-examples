FROM golang:1.15.1-alpine3.12 AS build-env

WORKDIR /tmp/simple-go-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM scratch
COPY --from=build-env /tmp/simple-go-app/simple-web-app /app/simple-web-app

EXPOSE 8080
CMD ["/app/simple-web-app"]
