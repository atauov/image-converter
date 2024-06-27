FROM golang:1.22.2-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add gcc make musl-dev bash wget && wget https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz && \
                                                  tar -zxvf migrate.linux-amd64.tar.gz && \
                                                  mv migrate /usr/local/bin/migrate && \
                                                  rm migrate.linux-amd64.tar.gz

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./bin/app cmd/image-converter/main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates make bash

COPY --from=builder /usr/local/src/bin/app /
COPY config/local.yaml /config/local.yaml
COPY .env /.env
COPY ./schema ./schema
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY Makefile /Makefile

EXPOSE 8889

LABEL maintainer="Aslan <atauov@gmail.com>"
LABEL version="1.0"
LABEL description="Test converter"