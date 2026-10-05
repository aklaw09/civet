# Civet

Civet is a small Redis-compatible server built in Go as a learning project.

## Run

```sh
go run ./cmd/civet -addr localhost
```

The server listens on `localhost:6730`. To verify the first implemented
behavior, connect from another terminal:

```sh
nc localhost 6730 | hexdump -C
```

The response is `+PONG\r\n`.

## Lesson 1

Lesson 1 creates a TCP listener that accepts a client connection, writes a
PONG response, and closes the connection.

## Not implemented yet

Civet does not yet parse Redis commands, store data, handle clients
concurrently, or persist data.
