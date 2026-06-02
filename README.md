# URL Checker

Concurrent URL checker written in Go.

Features:

- Fan-in, Fan-out
- Worker pool
- Context cancellation (Ctrl+C)
- Rate limiting
- Burst support
- Graceful shutdown

## Run

Check several URLs:

```bash
go run . \
https://google.com \
https://github.com \
https://golang.org \
https://example.com
```

## Example

```bash
go run . \
https://google.com \
https://github.com \
https://httpbin.org/status/404 \
https://httpbin.org/status/500
```

Example output:

```text
https://google.com -> 200
https://github.com -> 200
https://httpbin.org/status/404 -> 404
https://httpbin.org/status/500 -> 500
```

## Testing Rate Limiting

Run with multiple URLs:

```bash
go run . \
https://httpbin.org/delay/1 \
https://httpbin.org/delay/1 \
https://httpbin.org/delay/1 \
https://httpbin.org/delay/1 \
https://httpbin.org/delay/1 \
https://httpbin.org/delay/1 \
https://httpbin.org/delay/1
```

Observe that requests are released according to the configured rate limiter.

## Cancel Execution

Press:

```text
Ctrl+C
```

The application will stop processing and shut down gracefully.

## Stress Test

Run the checker with many URLs to observe:

- worker pool behavior
- rate limiting
- burst capacity
- graceful shutdown

```bash
go run . \
https://google.com \
https://github.com \
https://golang.org \
https://example.com \
https://google.com \
https://github.com \
https://golang.org \
https://example.com \
https://google.com \
https://github.com \
https://golang.org \
https://example.com \
https://google.com \
https://github.com \
https://golang.org \
https://example.com \
https://google.com \
https://github.com \
https://golang.org \
https://example.com
```
or 
```
go run . $(yes https://google.com | head -50)
```