package emitter

import (
	"fmt"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/field"
)

type Emitter interface {
	Emit(ctx *context.Context) string
}

// NewFunc is a function that creates a new Emitter with the given config.
type NewFunc = func(cfg *ucfg.Config, fields []*field.Field) (Emitter, error)

type config struct {
	Type string `config:"type" validate:"required"`
}

var registry = map[string]NewFunc{}

// New creates a new Emitter from the given config.
func New(cfg *ucfg.Config, fields []*field.Field) (Emitter, error) {
	c := config{}
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	newFn, exists := registry[c.Type]
	if !exists {
		return nil, fmt.Errorf("emitter: unknown emitter: %q", c.Type)
	}

	return newFn(cfg, fields)
}

// Register will register an emitter with a given name.
func Register(name string, fn NewFunc) {
	if _, exists := registry[name]; exists {
		panic("emitter: '" + name + "' already registered")
	}
	registry[name] = fn
}
