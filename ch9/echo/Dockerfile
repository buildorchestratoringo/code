FROM golang:1.17.8-alpine3.14 as builder
RUN mkdir /build
ADD *.go go.* /build/
WORKDIR /build
RUN go build .

FROM alpine:3.14
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/echo .

CMD /app/echo
