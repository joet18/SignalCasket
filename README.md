# SignalCasket

A real-time log aggregation engine built with **Go** and **gRPC streaming**, featuring an HTTP query API for efficient log retrieval and monitoring.

## Overview

SignalCasket is a lightweight, high-performance log tail and aggregation service. It monitors file changes in real-time, streams new log entries, and provides HTTP endpoints for querying and monitoring log data.

## Features

- **Real-time File Tailing**: Continuously monitors log files for new entries
- **Concurrent Log Aggregation**: Thread-safe log collection with mutex synchronization
- **HTTP Query API**: RESTful endpoints for retrieving logs with optional filtering
- **Graceful Shutdown**: Handles SIGINT and SIGTERM signals cleanly
- **Status Monitoring**: Built-in health check endpoint
- **gRPC-Ready**: Designed with streaming capabilities for distributed log collection

## Technology Stack

- **Language**: Go 1.26.4
- **Concurrency**: Goroutines and channels for efficient async processing
- **Protocols**: HTTP/REST and gRPC (streaming)
- **Synchronization**: sync.Mutex for thread-safe operations

## Project Structure

```
SignalCasket/
├── main.go         # Entry point, file tailer, and main event loop
├── handler.go      # HTTP request handlers for logs and status
├── go.mod          # Go module definition
├── test.log        # Sample log file for testing
└── README.md       # This file
```

## Getting Started

### Prerequisites

- Go 1.26.4 or later

### Installation

Clone the repository:

```bash
git clone https://github.com/joet18/SignalCasket.git
cd SignalCasket
```

### Building

```bash
go build -o signalcasket
```

### Running

Provide a log file path as an argument:

```bash
./signalcasket /path/to/your/logfile.log
```

Example:

```bash
./signalcasket test.log
```

## API Endpoints

### 1. **GET /logs**

Retrieve aggregated log entries.

**Parameters:**
- `limit` (optional): Maximum number of recent log lines to return

**Examples:**

```bash
# Get all logs
curl http://localhost:8080/logs

# Get last 10 logs
curl http://localhost:8080/logs?limit=10

# Get last 50 logs
curl http://localhost:8080/logs?limit=50
```

**Response:**

```json
[
  "2026-07-17 10:15:23 INFO Application started",
  "2026-07-17 10:15:24 DEBUG Connection established",
  "2026-07-17 10:15:25 ERROR Failed to connect to database"
]
```

### 2. **GET /status**

Health check endpoint to verify the service is running.

**Example:**

```bash
curl http://localhost:8080/status
```

**Response:**

```
Signalcasket is running
```

## Architecture

### Main Components

#### `main.go`
- **main()**: Initializes the application, sets up signal handling, starts HTTP server, and begins file tailing
- **tailFile()**: Continuously reads new lines from the log file with polling and graceful shutdown support
- **checkArgs()**: Validates command-line arguments

**Key Features:**
- Signal handling for clean shutdown (SIGINT, SIGTERM)
- Channel-based communication between goroutines
- Context-based cancellation propagation

#### `handler.go`
- **logsHandler()**: Processes `/logs` requests, supports limit-based pagination
- **statHandler()**: Simple status endpoint
- **Thread-safe access**: Uses `sync.Mutex` to protect shared log slice

## How It Works

1. **Startup**: User provides a log file path
2. **File Tailing**: Application opens the file and reads existing content
3. **Real-time Monitoring**: Continuously polls for new lines (1-second intervals)
4. **Log Aggregation**: New lines are pushed to a thread-safe slice via channel
5. **HTTP Server**: Listens on `:8080` for incoming requests
6. **Query Processing**: Clients can fetch logs with optional filtering
7. **Graceful Shutdown**: Responds to OS signals and cleanly exits

### Concurrency Model

```
┌─────────────────────────────────────────┐
│         Main Goroutine                  │
│  - Signal handling                      │
│  - Event loop (select on channels)      │
│  - Log aggregation                      │
└─────────────────────────────────────────┘
         ↓                    ↓
    ┌────────────┐      ┌────────────┐
    │  HTTP      │      │  File      │
    │  Server    │      │  Tailer    │
    │  Goroutine │      │  Goroutine │
    └────────────┘      └────────────┘
         ↓                    ↓
    ":8080/logs"         Read from
    ":8080/status"       log file
```

## Usage Examples

### Start SignalCasket

```bash
./signalcasket test.log
# Output:
# starting tailer..
# Tailing file: test.log
# waiting formare EOF
# Waiting
# ...
```

### Query Logs via HTTP

```bash
# Get all logs
curl http://localhost:8080/logs | jq

# Get last 5 entries
curl "http://localhost:8080/logs?limit=5" | jq

# Check status
curl http://localhost:8080/status
```

### Graceful Shutdown

Press `Ctrl+C` (SIGINT) to gracefully shut down:

```
^Cshutdown signal recived
Exiting..
```

## Configuration

Currently, SignalCasket uses default configurations:
- **HTTP Port**: 8080
- **File Poll Interval**: 1 second
- **Log Storage**: In-memory slice

Future enhancements may include:
- Configurable port and polling interval
- Persistent storage backends
- Log filtering and search capabilities
- Metrics and monitoring

## Current Limitations

- Single file tailing (one file per instance)
- In-memory log storage (no persistence)
- No built-in log rotation handling
- No authentication/authorization
- Basic error handling

## Roadmap

- [ ] Support for multiple file sources
- [ ] Persistent log storage (database backend)
- [ ] Advanced log filtering and search
- [ ] gRPC streaming implementation
- [ ] Metrics and observability (Prometheus)
- [ ] Configuration file support
- [ ] Docker containerization
- [ ] Authentication and authorization

## Contributing

Contributions are welcome! Feel free to:
- Report issues
- Suggest features
- Submit pull requests

## License

This project is open source. See LICENSE file for details (if applicable).

## Author

**joet18** - [GitHub Profile](https://github.com/joet18)

---

## Development Notes

- Code follows Go conventions and best practices
- Signal handling ensures graceful application termination
- Thread-safe operations via mutex protection
- Ready for enhancement with gRPC streaming capabilities
