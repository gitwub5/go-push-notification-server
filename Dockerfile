FROM golang:alpine AS builder  

LABEL maintainer="ssgwoo <tonyw2@khu.ac.kr>"

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod download && go build -o go-notification-server ./cmd/main.go

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/go-notification-server .

EXPOSE 8080

CMD ["./go-notification-server"]