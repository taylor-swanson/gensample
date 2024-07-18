# File Output (file)

The `file` output writes sample data to a file. The output can be configured to
output to a singular file or to directory with a random name.

## Configuration

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
