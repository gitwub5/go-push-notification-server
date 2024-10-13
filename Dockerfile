FROM golang:alpine AS builder  

LABEL maintainer="GeonWoo <tonyw2@khu.ac.kr>"

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod download && go build -o notification-server ./cmd/main.go

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/notification-server .

EXPOSE 8080

CMD ["./notification-server"]