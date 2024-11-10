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
   curl -X POST http://localhost:8080/send     
        -H "Content-Type: application/json"     
        -d '{
            "title": "Hello",
            "message": "This is a test",
            "token": "example-token",
            "priority": "high",
            "platform": 2
            }'
    ```

    **Example Response**:
    ```json
    {
        "status": "success",
        "data": {
            "notification_id": "cb7675bb-37ea-46af-a0c0-98004eeaff03",
            "title": "Hello",
            "message": "This is a test"
        }
    }
    ```

    **Note**: The `notification_id` for each notification is generated using a UUID to ensure global uniqueness.

3. Send a Subscription request & Unsubscription request:
    ```sh
    curl -X POST http://localhost:8080/subscribe \
        -H "Content-Type: application/json" \
        -d '{
        "token": "example-device-token",
        "topic": "primary notification",
        "platform": 2
        }'
    ```

    **Example Response**:
    ```json
    {
        "status": "success",
        "message": "Subscribed to topic successfully!"
    }
    ```

    ```sh
    curl -X POST http://localhost:8080/unsubscribe \
        -H "Content-Type: application/json" \
        -d '{
        "token": "example-device-token",
        "topic": "primary notification",
        "platform": 2
        }'
    ```

    **Example Response**:
    ```json
    {
        "status": "success",
        "message": "Unsubscribed from topic successfully!"
    }
    ```
    
## Added APIs

### 1. **Notification Status API**

**Endpoint**: `GET /api/status/{notification_id}`

This API allows you to check the status of a specific notification that has been sent. It returns whether the notification was successfully delivered or failed.

**Example**:
```sh
curl -X GET http://localhost:8080/api/status/1234f963-e9d5-488d-93cb-2fc83db02fcc
```

**Response**:
```json
{
    "notification_id": "1234f963-e9d5-488d-93cb-2fc83db02fcc",
    "title": "Test Notification",
    "message": "This is a test message",
    "priority": "high",
    "status": "delivered"
}
```

**Note**: The `notification_id` field represents a UUID generated for the notification, ensuring it is globally unique.

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
        "id": "73ee262d-c2ad-405c-b4b1-aeace2f8029f",
        "title": "Test Notification",
        "message": "This is a test message",
        "priority": "high",
        "status": "delivered"
    },
    {
        "id": "b5676a2d-8c25-4893-b8e2-abc1234ef567",
        "title": "Another Notification",
        "message": "Another test message",
        "priority": "normal",
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

Configuration options can be set in the `config.yml` file or overridden using environment variables. This allows flexibility for different deployment environments (e.g., development vs. production).

### Example `config.yml` File:

```yaml
server:
  port: 8080  # The port the server will run on

redis:
  host: localhost  # Redis host address
  port: 6379       # Redis port number
  password: yourpassword  # Redis password (leave empty if not set)

mysql:
  host: localhost  # MySQL host address
  port: 3306       # MySQL port number
  user: root       # MySQL user
  password: password  # MySQL password
  database: push_notification_db  # Database name
```

### Environment Variable Overrides:
The server configuration can be overridden using environment variables for better flexibility in deployment. Here are some examples of environment variables that can be used:

- `SERVER_PORT`: Overrides the server port.
- `REDIS_HOST`: Overrides the Redis host address.
- `REDIS_PORT`: Overrides the Redis port number.
- `REDIS_PASSWORD`: Overrides the Redis password.
- `MYSQL_HOST`: Overrides the MySQL host address.
- `MYSQL_PORT`: Overrides the MySQL port number.
- `MYSQL_USER`: Overrides the MySQL user.
- `MYSQL_PASSWORD`: Overrides the MySQL password.
- `MYSQL_DATABASE`: Overrides the MySQL database name.

**Note**: If environment variables are set, they will take precedence over the `config.yml` file.


## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, please contact [ssgwoo5@gmail.com](mailto:ssgwoo5@gmail.com).
