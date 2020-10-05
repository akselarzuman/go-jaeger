FROM alpine:3.12 as base
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 80

# Build the code
FROM golang:1.15.0-alpine3.12 as build

# Add git dependency for imported projects
RUN apk add git

WORKDIR /go/jaeger-test-api
COPY . /go/jaeger-test-api

# Build the Go app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/jaeger-test-api/main /go/jaeger-test-api/main.go

FROM base as final
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
# Change TimeZone
RUN apk add --update tzdata
ENV TZ=Asia/Istanbul
# Clean APK cache
RUN rm -rf /var/cache/apk/*
WORKDIR /root/
COPY --from=build /go/jaeger-test-api .

# Command to run the executable
CMD ["./main"]