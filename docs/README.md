# Docs

Docs for `gensample`.

- [Emitters](emitters.md)
- [Generators](generators.md)
- [Outputs](outputs.md)

## Configuration

- `emitter` The [emitter](emitters.md) configuration.
- `output` The [output](outputs.md) configuration.
- `fields` The fields to be used for data generation. Each field contains a [generator](generators.md) configuration.
- `interval`: How often to send records. Omit for one-shot.
- `records`: The number of records to send each interval.

```yaml
---
interval: 1s
records: 1
output:
  type: udp
  host: localhost
  port: 9001
  network: udp4
emitter:
  type: template
  templates:
    - '<140>{{.datetime}} myAsaHostname CiscoASA[999]: %ASA-6-305011: Built dynamic TCP translation from inside:{{.src_ip}}/{{.src_port}} to outside:{{.dst_ip}}/{{.dst_port}}'
    - '<140>{{.datetime}} myAsaHostname CiscoASA[999]: %ASA-6-302016: Teardown UDP connection {{.connection_id}} for outside:{{.dst_ip}}/{{.dst_port}} to inside:{{.src_ip}}/{{.src_port}} duration 0:00:00 bytes {{.byte_count}}'
fields:
  - name: datetime
    generator:
      type: date
      format: "Jan _2 2006 15:04:05"
  - name: src_ip
    generator:
      type: ip
      cidr: "10.0.0.0/24"
  - name: src_port
    generator:
      type: port
      min: 1024
      max: 16383
  - name: dst_port
    generator:
      type: port
      min: 16384
      max: 32767
  - name: dst_ip
    generator:
      type: ip
      cidr: "172.16.0.0/24"
  - name: connection_id
    generator:
      type: int
      min: 1000
      max: 10000
  - name: byte_count
    generator:
      type: int
      min: 50
      max: 150
```

### Fields

- `name` The name of the field.
- `generator` The [generator](generators.md) config for the field.

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
