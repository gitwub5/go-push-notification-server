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
    go run cmd/main.go
    ```
2. Send a test notification:
    ```sh
    curl -X POST http://localhost:8080/send     -H "Content-Type: application/json"     -d '{"title": "Hello", "message": "This is a test", "token": "example-token"}'
    ```
3. Send a Subscription request & Unsubscription request:
    ```sh
    curl -X POST http://localhost:8080/subscribe     -H "Content-Type: application/json"     -d '{
    "token": "example-device-token",
    "topic": "primary notification"
    }'
    ```

    ```sh
    curl -X POST http://localhost:8080/unsubscribe     -H "Content-Type: application/json"     -d '{
    "token": "example-device-token",
    "topic": "primary notification"
    }'
    ```

## Added APIs

### 1. **Notification Status API**

**Endpoint**: `GET /api/status/{notification_id}`

This API allows you to check the status of a specific notification that has been sent. It returns whether the notification was successfully delivered or failed.

**Example**:
```sh
curl -X GET http://localhost:8080/api/status/12345
```

**Response**:
```json
{
    "notification_id": "12345",
    "status": "delivered"
}
```

### 2. **Notification Logs API**

**Endpoint**: `GET /api/logs`

This API retrieves the logs of all notifications sent from the server, displaying information about the status of each notification.

**Example**:
```sh
curl -X GET http://localhost:8080/api/logs
```

**Response**:
```json
[
    {
        "notification_id": "12345",
        "title": "Test Notification",
        "status": "delivered"
    },
    {
        "notification_id": "67890",
        "title": "Another Notification",
        "status": "failed"
    }
]
```

### 3. **Health Check API**

**Endpoint**: `GET /api/health`

This API provides a quick way to check if the server is running and healthy. It will return a simple status message indicating whether the server is operational.

**Example**:
```sh
curl -X GET http://localhost:8080/api/health
```

**Response**:
```json
{
    "status": "healthy"
}
```

### 4. **Golang Performance Metrics API**

**Endpoint**: `GET /api/stat/go`

This API returns performance-related statistics about the Go runtime, such as memory usage, garbage collection, and CPU statistics.

**Example**:
```sh
curl -X GET http://localhost:8080/api/stat/go
```

**Response**:
```json
{
    "cpu": {
        "usage": "5%",
        "cores": 4
    },
    "memory": {
        "alloc": "20MB",
        "total": "100MB"
    },
    "gc": {
        "count": 10,
        "pause_total_ns": 200000
    }
}
```

### 5. **Notification Stats API**

**Endpoint**: `GET /api/stat/app`

This API provides application-level statistics about push notifications sent, such as how many notifications were successfully delivered and how many failed.

**Example**:
```sh
curl -X GET http://localhost:8080/api/stat/app
```

**Response**:
```json
{
    "success": 100,
    "failure": 5
}
```

### 6. **Server Configuration API**

**Endpoint**: `GET /api/config`

This API allows you to retrieve the current server configuration as set in the `config.yml` file. This is useful for debugging and ensuring that the server is using the correct configuration.

**Example**:
```sh
curl -X GET http://localhost:8080/api/config
```

**Response**:
```yaml
port: 8080
redis:
  host: localhost
  port: 6379
  password: yourpassword
```

## Configuration

Configuration options can be set in the `config.yml` file. Here is an example configuration:

```yaml
port: 8080
redis:
  host: localhost
  port: 6379
  password: yourpassword
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, please contact [ssgwoo5@gmail.com](mailto:ssgwoo5@gmail.com).
