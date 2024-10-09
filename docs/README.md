# Go Push Notification Server

This repository contains a Go-based push notification server. The server is designed to handle and send push notifications efficiently.

## Features

- **High Performance**: Optimized for handling a large number of push notifications.
- **Scalable**: Easily scalable to meet growing demands.
- **Secure**: Implements best practices for security.

## Requirements

- Go 1.16 or higher
- A working internet connection

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gitwub5/go-push-notification-server.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-push-notification-server
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Start the server:
    ```sh
    go run main.go
    ```
2. Send a test notification:
    ```sh
    curl -X POST http://localhost:8080/send \
    -H "Content-Type: application/json" \
    -d '{"title": "Hello", "message": "This is a test", "token": "example-token"}'
    ```
3. Send a Subscription request & Unsubscription request
     ```sh
    curl -X POST http://localhost:8080/subscribe \
    -H "Content-Type: application/json" \
    -d '{
    "token": "example-device-token",
    "topic": "primary notification"
    }'
    ```

     ```sh
    curl -X POST http://localhost:8080/unsubscribe \
    -H "Content-Type: application/json" \
    -d '{
    "token": "example-device-token",
    "topic": "primary notification"
    }'
    ```

## Configuration

Configuration options can be set in the `config.json` file. Here is an example configuration:

```json
{
    "port": 8080,
    "redis": {
        "host": "localhost",
        "port": 6379,
        "password": "yourpassword"
    }
}
```


## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, please contact [ssgwoo5@gmail.com](mailto:ssgwoo5@gmail.com).
