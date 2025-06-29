FROM golang:1.24-alpine AS build

WORKDIR /build
COPY ./go.mod /build/go.mod
COPY ./go.sum /build/go.sum

RUN go mod download && go mod verify

COPY . /build/

RUN go build -v

FROM alpine:3.21 AS run

WORKDIR /usr/bin
COPY --from=build /build/brightsky-temp .

RUN addgroup -S temperature
RUN adduser -S temperature temperature
USER temperature

ENTRYPOINT [ "/usr/bin/brightsky-temp" ]