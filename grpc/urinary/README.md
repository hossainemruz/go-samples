# Urinary RPC

Server receive single request, process it and send a single response.

## Code Generation

```bash
protoc api/greeter.proto \
  --go_out=gen \
  --go_opt=paths=source_relative \
  --go-grpc_out=gen \
  --go-grpc_opt=paths=source_relative
```

## Usage

**Run server:**

```bash
go run server/main.go --port=4321
```

**Run client:**

```bash
go run client/main.go --addr="localhost:4321" --name="John"
```
