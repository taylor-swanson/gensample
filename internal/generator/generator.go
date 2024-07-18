package generator

import (
	"fmt"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
)

type Generator interface {
	Generate(ctx *context.Context) string
}

// NewFunc is a function that creates a new Generator with the given config.
type NewFunc = func(*ucfg.Config) (Generator, error)

type config struct {
	Type string `config:"type" validate:"required"`
}

var registry = map[string]NewFunc{}

// New creates a new Generator from the given config.
func New(cfg *ucfg.Config) (Generator, error) {
	c := config{}
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	newFn, exists := registry[c.Type]
	if !exists {
		return nil, fmt.Errorf("generator: unknown generator: %q", c.Type)
	}

	return newFn(cfg)
}

// Register will register a generator with a given name.
func Register(name string, fn NewFunc) {
	if _, exists := registry[name]; exists {
		panic("generator: '" + name + "' already registered")
	}
	registry[name] = fn
}
