# Demo: gRPC reconnection ability

This is a example showing the gRPC reconnectability with UDS

## Testing

```
docker compose up
```

## Set up

- Server: a simple echo server that start a UDS. It crashes intentionally after 10 seconds to test the 
  reconnectability.
- Server: it deletes the UDS file when it starts.
- Client: Cron gRPC client through UDS.

