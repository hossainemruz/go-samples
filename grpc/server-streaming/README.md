# Server Streaming

Server receive a single request from client but send a stream of response.

## Example: Real-Time Stock Price Updates

Imagine we are building a financial app that require to give user real time stock price update. Instead of client fetching stock price periodically, we can use server streaming here. Server will push updated price as soon as they are available.

### Benefits

- **Real-Time Updates:** The client receives updates as soon as they are available, without having to poll the server repeatedly.
- **Efficiency:** Reduces network traffic and server load compared to polling.
- **Scalability:** The server can handle many concurrent clients efficiently by streaming updates to each client individually.

### Key Takeaways

- Server streaming is ideal for scenarios where the server needs to push a continuous stream of data to the client.
- The client receives updates as they become available, providing a real-time experience.
- This pattern can significantly improve efficiency and scalability compared to traditional request-response patterns.

## Code Generation

```bash
protoc api/stock_service.proto \
  --go_out=gen \
  --go_opt=paths=source_relative \
  --go-grpc_out=gen \
  --go-grpc_opt=paths=source_relative
```

## Run Server

```bash
go run server/main.go --port=4321
```

## Run Client

```bash
go run client/main.go --addr="localhost:4321" --symbol="NVDIA"
```
