# Date Generators

## Date (date)

Date and time generator. The value produced is based on the current time and
is not random.

### Configuration

- `format`: The date format string. Accepts [Go time format strings](https://go.dev/src/time/format.go), or one of the predefined values below:
    - `unix`: Unix timestamp, in seconds
    - `unix_ms`: Unix timestamp, in milliseconds
    - `unix_us`: Unix timestamp, in microseconds
    - `unix_ns`: Unix timestamp, in nanoseconds

### Usage

```yaml
type: date
format: "Jan 02 2006 15:04:05"
```

```yaml
type: date
format: "unix_ns"
```
