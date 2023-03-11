FROM golang:1.14.6-alpine3.12 as builder
WORKDIR /app
RUN go mod download
COPY . .
RUN 

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/gitlab.com/idoko/bucketeer/build/bucketeer /usr/bin/bucketeer
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/bucketeer"]