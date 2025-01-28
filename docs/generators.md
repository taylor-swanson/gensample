# Generators

## Date (date)

Date and time generator. The value produced is based on the current time and
is not random.

### Configuration

- `format`: The date format string. Accepts [Go time format strings](https://go.dev/src/time/format.go), or one of the predefined values below:
    - `rfc3339`: RFC 3339 timestamp
    - `unix`: Unix timestamp, in seconds
    - `unix_ms`: Unix timestamp, in milliseconds
    - `unix_us`: Unix timestamp, in microseconds
    - `unix_ns`: Unix timestamp, in nanoseconds
    - `unix_s_ms`: Unix timestamp in seconds with fractional microseconds

```yaml
type: date
format: "Jan 02 2006 15:04:05"
```

```yaml
type: date
format: "unix_ns"
```

## Country (country)

Country name generator.

### Configuration

- `values`: Override the default list with custom values.

```yaml
type: country
```

```yaml
type: country
values:
  - Australia
  - Canada
  - United States
```

## First Name (first_name)

First name generator.

### Configuration

- `values`: Override the default list with custom values.

```yaml
type: first_name
```

```yaml
type: first_name
values:
  - Alice
  - Bob
  - Chuck
```

## Last Name (last_name)

Last (surname) name generator.

### Configuration

- `values`: Override the default list with custom values.

```yaml
type: last_name
```

```yaml
type: last_name
values:
  - Anderson
  - Johnson
  - Smith
```

## Email (email)

Email generator. Format is `user@domain.tld`, where `user` and `domain` are
random words and `tld` is one of `com`, `net`, `org`.

### Configuration

- `values`: Override the default list with custom values.

```yaml
type: email
```

```yaml
type: email
values:
  - 'alice@example.com'
  - 'bob@example.net'
  - 'chuck@example.org'
```

## Port (port)

Port number generator.

### Configuration

- `min`: The minimum value.
- `min`: The maximum value.

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

```yaml
type: ip
cidr: '10.0.0.0/24'
```

## MAC Address (mac)

MAC address generator.

### Configuration

- `separator`: The separator to use between octets. Defaults to `-`.
- `lowercase`: Lowercase formatting. Defaults to `false`.

```yaml
type: mac
```

```yaml
type: mac
separator: ':'
lowercase: true
```

## Float (float)

Floating point value generator. Generates a random number between `min` and `max`.

### Configuration

- `min`: The minimum value to generate.
- `max`: The maximum value to generate.

```yaml
type: float
min: 1.5
max: 3.5
```

## Integer (int)

Integer value generator. Generates a random integer between `min` and `max`.

### Configuration

- `min`: The minimum value to generate.
- `max`: The maximum value to generate.

```yaml
type: date
min: 0
max: 10
```

## String (string)

Generate a string of random letters and numbers of specified length.

### Configuration

- `length`: The length of the string. Must be a positive integer.

```yaml
type: string
length: 8
```

## String Array (string_array)

Select a random string from a list.

### Configuration

- `strings`: The list of strings.

```yaml
type: string_array
strings:
  - alpha
  - beta
  - gamma
```

## Static String (static)

A fixed string value.

### Configuration

- `value`: The string value.

```yaml
type: static
value: "foobar"
```

## HTTP Status (http_status)

Generate a random HTTP status.

### Configuration

There is no configuration for the `http_status` generator.

```yaml
type: http_status
```

## HTTP Method (http_method)

Generate a random HTTP method.

### Configuration

There is no configuration for the `http_method` generator.

```yaml
type: http_method
```

## HTTP Version (http_version)

Generate a random HTTP version.

### Configuration

There is no configuration for the `http_version` generator.

```yaml
type: http_version
```

## User Agent (user_agent)

Generate a random user agent string.

### Configuration

There is no configuration for the `user_agent` generator.

```yaml
type: user_agent
```

## UUID (uuid)

UUID generator.

### Configuration

There is no configuration for the `uuid` generator.
