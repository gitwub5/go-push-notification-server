FROM golang:1.17

LABEL maintainer="ssgwoo <tonyw2@khu.ac.kr>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-push-client ./cmd/main.go

CMD ["./go-push-client"]