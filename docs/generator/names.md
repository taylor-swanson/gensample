# Name Generators

## Country (country)

Country name generator.

### Configuration

- `values`: Override the default list with custom values.

### Usage

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

### Usage

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

### Usage

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

### Usage

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
