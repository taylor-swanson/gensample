package uuid

import (
	"github.com/elastic/go-ucfg"
	"github.com/google/uuid"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/generator"
)

const (
	NameUUID = "uuid"
)

type id struct{}

func (g *id) Generate(ctx *context.Context) string {
	v, _ := uuid.NewRandomFromReader(&generator.RandReader{Rand: ctx.Rand})

	return v.String()
}

func NewUUIDGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := id{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

func init() {
	generator.Register(NameUUID, NewUUIDGenerator)
}
