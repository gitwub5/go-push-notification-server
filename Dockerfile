FROM golang:alpine AS builder  

LABEL maintainer="ssgwoo <tonyw2@khu.ac.kr>"

WORKDIR /app

ENV CGO_ENABLED=1
RUN apk add --no-cache git gcc musl-dev

COPY . .

RUN go mod download && go build -o go-notification-server ./cmd/main.go

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/go-notification-server .

COPY --from=builder /app/config /root/config

EXPOSE 8080

CMD ["./go-notification-server"]