# Template Emitter (template)

## Configuration

- `template`: The template string.

## Usage

```yaml
---
emitter:
  type: template
  template: '<140>{{.datetime}} myAsaHostname: %ASA-6-716059: Group {{.group}} User {{.user}} IP {{.ip1}} AnyConnect session resumed. Connection from {{.ip2}}.'
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
    - name: user
      generator:
        type: string_array
        strings:
          - user_1
          - user_2
          - user_3
          - user_4
          - user_5
    - name: ip1
      generator:
        type: ip
        cidr: "10.0.0.0/24"
    - name: ip2
      generator:
        type: ip
        cidr: "172.16.0.0/24"
```

