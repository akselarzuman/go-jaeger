FROM golang:1.23-alpine3.21 as build

WORKDIR /go/jaeger-test-api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main ./api/main.go

FROM alpine:3.21 as final
WORKDIR /root/
COPY --from=build /go/jaeger-test-api/main .

EXPOSE 8080
ENTRYPOINT ["./main"]