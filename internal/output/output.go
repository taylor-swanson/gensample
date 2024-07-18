package output

import (
	"fmt"
	"io"

	"github.com/elastic/go-ucfg"
)

type Output interface {
	io.WriteCloser
	NewInterval() error
}

// NewFunc is a function that creates a new Output with the given config.
type NewFunc = func(*ucfg.Config) (Output, error)

type config struct {
	Type string `config:"type" validate:"required"`
}

var registry = map[string]NewFunc{}

// New creates a new Output from the given config.
func New(cfg *ucfg.Config) (Output, error) {
	c := config{}
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	newFn, exists := registry[c.Type]
	if !exists {
		return nil, fmt.Errorf("output: unknown output: %q", c.Type)
	}

	return newFn(cfg)
}

// Register will register an output with a given name.
func Register(name string, fn NewFunc) {
	if _, exists := registry[name]; exists {
		panic("output: '" + name + "' already registered")
	}
	registry[name] = fn
}
