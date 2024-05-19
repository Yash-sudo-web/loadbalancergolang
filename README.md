# Simple Load Balancer

This project implements a basic load balancer in Go, using the round-robin algorithm to distribute incoming HTTP requests to multiple backend servers.

## Features

- **Round-Robin Load Balancing**: Distributes requests evenly across the backend servers.
- **Health Check**: Basic health check to ensure servers are alive before forwarding requests (currently, all servers are assumed to be alive).
- **Reverse Proxy**: Uses Go's `httputil.ReverseProxy` to forward requests to backend servers.

## Getting Started

### Prerequisites

- Go (1.15 or higher)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/simple-load-balancer.git
    cd simple-load-balancer
    ```

2. Build the project:

    ```sh
    go build
    ```

3. Run the load balancer:

    ```sh
    ./simple-load-balancer
    ```

## Configuration

Currently, the backend servers are hardcoded in the `main` function. You can modify the list of servers by editing the following lines in `main.go`:

```go
servers := []Server{
    newSimpleServer("http://youtube.com"),
    newSimpleServer("http://facebook.com"),
    newSimpleServer("http://google.com"),
}
```

Replace the URLs with the addresses of your backend servers.

### Usage

After starting the load balancer, it listens on port 8080 (or the port specified in the main function). Incoming requests to localhost:8080 will be distributed to the backend servers in a round-robin fashion.

To test the load balancer, you can use curl or a web browser:

```sh
curl http://localhost:8080
```

### Contributing
Feel free to submit issues and pull requests to improve the functionality and reliability of the load balancer.
