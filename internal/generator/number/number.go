package number

import (
	"fmt"
	"math"
	"strconv"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/generator"
)

const (
	NameFloat = "float"
	NameInt   = "int"
)

type FloatGenerator struct {
	Min float64 `config:"min" validate:"required"`
	Max float64 `config:"max" validate:"required"`
}

func (g *FloatGenerator) Validate() error {
	if g.Min > g.Max {
		return fmt.Errorf("float: invalid range, min > max (%f > %f)", g.Min, g.Max)
	}

	return nil
}

func (g *FloatGenerator) Generate(ctx *context.Context) string {
	return fmt.Sprintf("%f", math.Mod(ctx.Rand.Float64(), g.Max-g.Min)+g.Min)
}

func NewFloatGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := FloatGenerator{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

type IntGenerator struct {
	Min int `config:"min"`
	Max int `config:"max"`
}

func (g *IntGenerator) Validate() error {
	if g.Min > g.Max {
		return fmt.Errorf("float: invalid range, min > max (%d > %d)", g.Min, g.Max)
	}

	return nil
}

func (g *IntGenerator) Generate(ctx *context.Context) string {
	return strconv.Itoa(ctx.Rand.IntN(g.Max-g.Min) + g.Min)
}

func NewIntGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := IntGenerator{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

func init() {
	generator.Register(NameFloat, NewFloatGenerator)
	generator.Register(NameInt, NewIntGenerator)
}
