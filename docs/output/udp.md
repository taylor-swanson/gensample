# UDP Output (udp)

The `udp` output writes sample data to a UDP socket.

## Configuration

- `host`: The remote host.
- `port`: The remote port.
- `network`: The network. Accepted values are `udp`, `udp4`, and `udp6`. Defaults to `udp`.

```yaml
output:
  type: udp
  host: 10.0.0.1
  port: 9001
```
