# Civet

Civet is a small Redis-compatible server built in Go as a learning project.

## Run

```sh
go run ./cmd/civet -addr 127.0.0.1:6380
```

The server listens on `127.0.0.1:6380` by default. To verify the first
implemented behavior, connect from another terminal:

```sh
nc 127.0.0.1 6380 | hexdump -C
```

The response is `+PONG\r\n`.

## Lesson 1

Lesson 1 creates a TCP listener that accepts a client connection, writes a
PONG response, and closes the connection.

## Not implemented yet

Civet does not yet parse Redis commands, store data, handle clients
concurrently, or persist data.
