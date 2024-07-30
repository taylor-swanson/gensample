# Outputs

## File Output (file)

The `file` output writes sample data to a file. The output can be configured to
output to a singular file or to directory with a random name.

### Configuration

- `filename`: The filename. If set, overrides the `directory` and `pattern` variables.
- `directory`: The output directory for writing files with random names.
- `pattern`: The pattern to use when writing files with random names.
- `delimiter`: The string to write between sample data. Defaults to `\n`.

```yaml
output:
  type: file
  filename: "/tmp/sample.log"
```

```yaml
output:
  type: file
  directory: "/tmp"
  pattern: "sample_*.log"
```

## stdout Output

The `stdout` output writes sample data to stdout.

### Configuration

There is no configuration for the `stdout` output.

```yaml
output:
  type: stdout
```

## TCP Output (tcp)

The `tcp` output writes sample data to a TCP stream.

### Configuration

- `host`: The remote host.
- `port`: The remote port.
- `network`: The network type. Accepted values are `tcp`, `tcp4`, and `tcp6`. Defaults to `tcp`.
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

## UDP Output (udp)

The `udp` output writes sample data to a UDP socket.

### Configuration

- `host`: The remote host.
- `port`: The remote port.
- `network`: The network type. Accepted values are `udp`, `udp4`, and `udp6`. Defaults to `udp`.

```yaml
output:
  type: udp
  host: 10.0.0.1
  port: 9001
```
