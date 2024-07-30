# gensample

Sample data generator.

## Building

The provided magefile can be used to build the project.

```shell
mage build
```

### Manual Build

```shell
go build -o build/gensample github.com/taylor-swanson/gensample/cmd/gensample
```

## Usage

1. Create a config for generating sample data. See `examples/` for example configurations.
2. Run the tool

```shell
build/gensample -c examples/demo.yml
```

## Documentation

- See [docs](docs/README.md) for documentation.
