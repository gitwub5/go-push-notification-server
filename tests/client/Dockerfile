FROM golang:1.23 AS builder

LABEL maintainer="ssgwoo <tonyw2@khu.ac.kr>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-push-client ./cmd/main.go

FROM alpine:3.18

WORKDIR /root/

RUN apk add --no-cache curl

COPY --from=builder /app/go-push-client .
COPY --from=builder /app/templates /root/templates

RUN chmod +x ./go-push-client

ENV TEMPLATE_PATH="/root/templates/*"

EXPOSE 8081

# 애플리케이션 실행
CMD ["./go-push-client"]