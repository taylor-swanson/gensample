package strings

import (
	"errors"
	"github.com/elastic/go-ucfg"

	"gensample/internal/context"
	"gensample/internal/generator"
)

const (
	NameString      = "string"
	NameStringArray = "string_array"
	NameStatic      = "static"
)

type str struct {
	Format string `config:"format" validate:"required"`
}

func (g *str) Generate(_ *context.Context) string {
	return "" // TODO!
}

func NewStringGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := str{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

type stringArray struct {
	Strings []string `config:"strings" validate:"required"`
}

func (g *stringArray) Validate() error {
	if len(g.Strings) == 0 {
		return errors.New("string_array: requires at least one string")
	}

	return nil
}

func (g *stringArray) Generate(ctx *context.Context) string {
	return g.Strings[ctx.Rand.IntN(len(g.Strings))]
}

func NewStringArrayGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := stringArray{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

type static struct {
	Value string `config:"value" validate:"required"`
}

func (g *static) Generate(_ *context.Context) string {
	return g.Value
}

func NewStaticGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := static{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

func init() {
	generator.Register(NameString, NewStringGenerator)
	generator.Register(NameStringArray, NewStringArrayGenerator)
	generator.Register(NameStatic, NewStaticGenerator)
}
