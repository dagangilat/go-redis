# Go-Redis Key-Value Store
This project implements a simple key-value store with a REST API, using Go and Redis.

## Features
- Set key-value pairs
- Retrieve values by key

## Prerequisites
- [Go](https://golang.org/dl/) (1.16 or later)
- [Redis](https://redis.io/download) (6.0 or later)

## Getting Started
1. Clone the repository:
    ```bash
    git clone https://github.com/dagangilat/go-redis.git
    cd go-redis
    ```

2. Install the required Go packages:
    ```bash
    go mod tidy
    ```

3. Start the Redis server (you may need to adjust this command depending on your Redis installation):
    ```bash
    redis-server
    ```

4. Run the Go server:
    ```bash
    go run main.go
    ```

The server will start on `localhost:8000`.

## API Usage

The API has two endpoints:

- `POST /set`
- `GET /get/{key}`

### Set a Key-Value Pair

Use the `/set` endpoint to set a key-value pair. The key-value pair should be provided in the request body as form data.

Here is an example using `curl`:

```bash
curl -X POST -d "key1=value1" http://localhost:8000/set
```

### Retrieve a Value by Key

Use the `/get/{key}` endpoint to retrieve a value by key. Replace `{key}` with the key of the value you want to retrieve.

Here is an example using `curl`:

```bash
curl http://localhost:8000/get/key1
```
