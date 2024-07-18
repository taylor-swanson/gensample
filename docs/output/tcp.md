# TCP Output (tcp)

The `tcp` output writes sample data to a TCP stream.

## Configuration

- `host`: The remote host.
- `port`: The remote port.
- `network`: The network. Accepted values are `tcp`, `tcp4`, and `tcp6`. Defaults to `tcp`.
- `octet-framing`: If true, enable octet framing. Defaults to `false`.
- `delimiter`: When non-transparent framing is used, the character used as a delimiter. Defaults to `\n`.

```yaml
output:
  type: tcp
  host: 10.0.0.1
  port: 9001
```

```yaml
output:
  type: tcp
  host: 10.0.0.1
  port: 9001
  network: tcp4
  octet-framing: true
```
