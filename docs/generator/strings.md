# String Generators

## String (string)

Not yet implemented.

## String Array (string_array)

Select a random string from a list.

### Configuration

- `strings`: The list of strings.

### Usage

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

### Usage

```yaml
type: static
value: "foobar"
```
