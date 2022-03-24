FROM golang:1.18.0-alpine3.15 as build

WORKDIR /go/jaeger-test-api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main ./api/main.go

FROM alpine:3.15.2 as final
WORKDIR /root/
COPY --from=build /go/jaeger-test-api/main .

EXPOSE 8080
ENTRYPOINT ["./main"]