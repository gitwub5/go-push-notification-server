# Go Push Client

This project is a Go-based client application that interacts with a push notification server. The client sends various requests to the server, including sending notifications, subscribing/unsubscribing from notifications, checking server health, and retrieving logs and metrics.

## Project Structure
```
go-push-client/
│
├── api
│   └── client.go           # API request handling logic
│
├── cmd
│   └── main.go             # Main entry point for the application
│
├── config
│   ├── config.go           # Configuration struct and environment variable parsing
│   └── config.yml          # Configuration file (example)
│
├── utils
│   └── logger.go           # Utility for logging
│
├── Dockerfile              # Dockerfile for building the Go client image
├── go.mod                  # Go module file
├── go.sum                  # Go module dependencies
└── README.md               # Project documentation
```

## How to Run the Project
### Prerequisites
- Go 1.16 or higher
- Docker (for containerized builds and runs)

### Running Locally
1. Clone the repository:
   ```bash
   git clone https://github.com/ssgwoo/go-push-client.git
   cd go-push-client
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run cmd/main.go
   ```

### Running with Docker
1. Build the Docker image:
   ```bash
   docker build -t ssgwoo/go-push-client .
   ```

2. Run the Docker container:
   ```bash
   docker run --rm -it ssgwoo/go-push-client
   ```

## API Requests in Main
The application sends the following API requests to the server:

1. **Send a test notification**
2. **Subscribe to notifications**
3. **Unsubscribe from notifications**
4. **Check the status of a notification**
5. **Get all notification logs**
6. **Check server health**
7. **Get Go runtime performance metrics**
8. **Get app-level notification statistics**
9. **Get server configuration**

Each request is sent with a 3-second interval between them for clarity.

## Sample Code Snippet
```go
package main

import (
    "fmt"
    "os"
    "time"

    "github.com/gitwub5/go-push-client/api"
)

func main() {
    baseURL := os.Getenv("SERVER_URL")
    if baseURL == "" {
        baseURL = "http://localhost:8080" // Default URL
    }
    fmt.Printf("Base URL: %s\n", baseURL)

    // Send a test notification
    fmt.Println("1. Sending a test notification...")
    notification := api.Notification{
        Title:    "Hello",
        Message:  "This is a test notification",
        Token:    "example-token",
        Priority: "high",  // Added priority
        Platform: 2,       // Added platform (e.g., 1 = iOS, 2 = Android)
    }
    notificationID, err := api.SendNotification(baseURL, notification)
    if err != nil {
        fmt.Printf("Failed to send notification: %v\n", err)
        return
    }
    time.Sleep(3 * time.Second) // 3-second delay

    // Subscribe to notifications
    fmt.Println("2. Subscribing client to notifications...")
    subRequest := api.SubscribeRequest{
        Token: "example-device-token",
        Topic: "primary notification",
    }
    api.Subscribe(baseURL, subRequest)
    time.Sleep(3 * time.Second) // 3-second delay

    // Unsubscribe from notifications
    fmt.Println("3. Unsubscribing client from notifications...")
    api.Unsubscribe(baseURL, subRequest)
    time.Sleep(3 * time.Second) // 3-second delay

    // Check the status of a notification
    fmt.Println("4. Checking the status of a notification...")
    if notificationID != "" {
        api.CheckNotificationStatus(baseURL, notificationID)
    } else {
        fmt.Println("No notification ID available to check status.")
    }
    time.Sleep(3 * time.Second) // 3-second delay

    // Get all notification logs
    fmt.Println("5. Getting all notification logs...")
    api.GetNotificationLogs(baseURL)
    time.Sleep(3 * time.Second) // 3-second delay

    // Check server health
    fmt.Println("6. Checking server health...")
    api.CheckServerHealth(baseURL)
    time.Sleep(3 * time.Second) // 3-second delay

    // Get Go runtime performance metrics
    fmt.Println("7. Getting Go runtime performance metrics...")
    api.GetGoPerformanceMetrics(baseURL)
    time.Sleep(3 * time.Second) // 3-second delay

    // Get app-level notification statistics
    fmt.Println("8. Getting app-level notification statistics...")
    api.GetNotificationStats(baseURL)
    time.Sleep(3 * time.Second) // 3-second delay

    // Get server configuration
    fmt.Println("9. Getting server configuration...")
    api.GetServerConfig(baseURL)

    // Print execution completion message
    fmt.Println("Execution has completed.")
}
```

## Maintainer
This project is maintained by **ssgwoo** (<tonyw2@khu.ac.kr>).
