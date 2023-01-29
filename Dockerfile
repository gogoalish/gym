
# Forum
FROM golang:1.19.4-alpine3.17 as forum

WORKDIR /app

COPY . .

RUN make build

# Distribution
FROM alpine:latest

WORKDIR /app 

EXPOSE 8080

COPY --from=forum /app/forum /app/

CMD /app/forum