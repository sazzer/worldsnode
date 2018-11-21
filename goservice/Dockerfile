FROM golang:1.11.2-alpine3.8 AS build

ENV GO111MODULE=on
WORKDIR /go/src/worlds
COPY . .
RUN apk update && apk upgrade && apk add --no-cache git gcc musl-dev
RUN go mod tidy
RUN go build -o worlds cmd/service/main.go

FROM alpine:3.8
WORKDIR /opt/worlds
COPY --from=build /go/src/worlds/worlds /opt/worlds/worlds
CMD /opt/worlds/worlds
