FROM golang:1.14.9 AS builder
WORKDIR /src
COPY . /src
RUN CGO_ENABLED=0 go build -o /go/bin/go-web

FROM ineva/alpine:3.9
WORKDIR /app
COPY --from=builder /go/bin/go-web /app/go-web
COPY /html /html
ENTRYPOINT /app/go-web -p 8080 -d /html

