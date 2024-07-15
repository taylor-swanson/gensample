# Network Generators

## Port (port)

Port number generator.

### Configuration

- `min`: The minimum value.
- `min`: The maximum value.

### Usage

```yaml
type: port
min: 443
max: 443
```

```yaml
type: port
min: 1024
max: 65535
```

## IP Address (ip)

IP address generator.

### Configuration

- `cidr`: The CIDR range.

### Usage

```yaml
type: ip
cidr: '10.0.0.0/24'
```

## MAC Address (mac)

MAC address generator.

### Configuration

- `separator`: The separator to use between octets. Defaults to `-`.
- `lowercase`: Lowercase formatting. Defaults to `false`.

### Usage

```yaml
type: mac
```

```yaml
type: mac
separator: ':'
lowercase: true
```
