// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package date

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/generator"
)

const (
	NameDate        = "date"
	NameBoundedDate = "bounded_date"
)

type date struct {
	Layout   string `config:"format" validate:"required"`
	TimeZone string `config:"timezone"`
}

func (g *date) Generate(_ *context.Context) string {
	now := time.Now()

	return formatTime(now, g.Layout)
}

func NewDateGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := date{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

type boundedDate struct {
	Start    string           `config:"start" validate:"required"`
	End      string           `config:"end" validate:"required"`
	Layout   string           `config:"format" validate:"required"`
	TimeZone string           `config:"timezone"`
	Easing   generator.Easing `config:"easing"`

	from float64
	to   float64
}

func (g *boundedDate) Generate(ctx *context.Context) string {
	t := g.Easing.Value(ctx.Progress)
	v := int64(g.from + (g.to-g.from)*t)
	tv := time.Unix(v, 0)

	return formatTime(tv, g.Layout)
}

func NewBoundedDateGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := boundedDate{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	from, err := time.Parse(time.RFC3339, g.Start)
	if err != nil {
		return nil, fmt.Errorf("generator.bounded_date: invalid start time: %w", err)
	}
	to, err := time.Parse(time.RFC3339, g.End)
	if err != nil {
		return nil, fmt.Errorf("generator.bounded_date: invalid end time: %w", err)
	}

	g.from = float64(from.Unix())
	g.to = float64(to.Unix())

	return &g, nil
}

func init() {
	generator.Register(NameDate, NewDateGenerator)
	generator.Register(NameBoundedDate, NewBoundedDateGenerator)
}

func formatTime(t time.Time, layout string) string {
	if layout == "rfc3339" {
		return t.Format(time.RFC3339)
	}
	if layout == "unix" {
		return strconv.FormatInt(t.Unix(), 10)
	}
	if layout == "unix_ms" {
		return strconv.FormatInt(t.UnixMilli(), 10)
	}
	if layout == "unix_us" {
		return strconv.FormatInt(t.UnixMicro(), 10)
	}
	if layout == "unix_ns" {
		return strconv.FormatInt(t.UnixNano(), 10)
	}
	if layout == "unix_s_us" {
		seconds := t.Unix()
		fractional := t.Nanosecond() / 1000

		return fmt.Sprintf("%d.%d", seconds, fractional)
	}

	return t.Format(layout)
}
