FROM golang:1.7.1 AS build-env
COPY src /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/sample src/sample/trivial-web-server.go

FROM scratch
COPY --from=build-env /go/bin/sample /app/sample

EXPOSE 8080
CMD ["/app/sample"]