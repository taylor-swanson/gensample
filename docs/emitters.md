# Emitters

Emitters are used to format the actual shape and data of a log.

## Template (template)

The Template emitter uses Go templates to produce sample data. The context for
the template is a map, where each key is the name of a `field`. Multiple templates
may be given, they will be advanced using a round-robin method during the
running of the tool.

### Configuration

- `templates`: A list of template strings.

```yaml
---
type: template
templates:
  - '<140>{{.datetime}} myAsaHostname CiscoASA[999]: %ASA-6-305011: Built dynamic TCP translation from inside:{{.src_ip}}/{{.src_port}} to outside:{{.dst_ip}}/{{.dst_port}}'
  - '<140>{{.datetime}} myAsaHostname CiscoASA[999]: %ASA-6-302016: Teardown UDP connection {{.connection_id}} for outside:{{.dst_ip}}/{{.dst_port}} to inside:{{.src_ip}}/{{.src_port}} duration 0:00:00 bytes {{.byte_count}}'
```
