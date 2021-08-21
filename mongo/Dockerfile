#stage 1
FROM golang:1.14-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
#download dependency
RUN go mod download
#copy source code
COPY . .
#build
RUN go build -o /go/bin/main

#stage 2
FROM alpine:3.5
WORKDIR /app
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

COPY --from=builder go/bin/main .  

#expose port
EXPOSE 9090
ENTRYPOINT ["./main"]