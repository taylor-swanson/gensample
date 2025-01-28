package strings

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/generator"
)

const (
	NameString      = "string"
	NameStringArray = "string_array"
	NameStatic      = "static"
)

var alphanum = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var hex = []rune("0123456789abcdef")

type str struct {
	Length int    `config:"length" validate:"required"`
	Format string `config:"format" validate:"required"`
}

func (g *str) Validate() error {
	if g.Length <= 0 {
		return fmt.Errorf("str: invalid length (%d), must be greater than zero", g.Length)
	}

	switch g.Format {
	case "alphanum", "hex":
		break
	default:
		return fmt.Errorf("str: invalid format: %s, must be one of 'alphanum', 'hex'", g.Format)
	}

	return nil
}

func (g *str) Generate(_ *context.Context) string {
	b := make([]rune, g.Length)
	for i := range b {
		switch g.Format {
		case "alphanum":
			b[i] = alphanum[rand.Intn(len(alphanum))]
		case "hex":
			b[i] = hex[rand.Intn(len(hex))]
		}
	}

	return string(b)
}

func NewStringGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := str{Format: "alphanum"}
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
