# Docs

Docs for `gensample`.

- [Emitters](emitter/README.md)
- [Generators](generator/README.md)

## Configuration

- `emitter` The [emitter](emitter/README.md) configuration.
- `fields` The fields.

### Fields

- `name` The name of the field.
- `generator` The [generator](generator/README.md) config for the field.

```yaml
fields:
  - name: datetime
    generator:
      type: date
      format: "Jan 02 2006 15:04:05"
  - name: group
    generator:
      type: string_array
      strings:
        - group_1
        - group_2
        - group_3
```
