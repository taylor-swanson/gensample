// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package field

import (
	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/generator"
)

type Field struct {
	Name        string
	AlwaysQuote bool
	Generator   generator.Generator
}

type config struct {
	Name        string       `config:"name" validate:"required"`
	AlwaysQuote bool         `config:"always_quote"`
	Generator   *ucfg.Config `config:"generator" validate:"required"`
}

func New(cfg *ucfg.Config) (*Field, error) {
	var c config
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	field := Field{
		Name:        c.Name,
		AlwaysQuote: c.AlwaysQuote,
	}
	var err error

	if field.Generator, err = generator.New(c.Generator); err != nil {
		return nil, err
	}

	return &field, nil
}
