// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package json

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/context"
	"github.com/taylor-swanson/gensample/internal/emitter"
	"github.com/taylor-swanson/gensample/internal/field"
)

const (
	Name = "json"
)

type config struct {
	Pretty bool `json:"pretty"`
}

type emt struct {
	fields  []*field.Field
	buf     *bytes.Buffer
	encoder *json.Encoder
}

func (e *emt) Emit(ctx *context.Context) string {
	values := make(map[string]string, len(e.fields))
	for _, v := range e.fields {
		values[v.Name] = v.Generator.Generate(ctx)
	}

	if err := e.encoder.Encode(values); err != nil {
		panic(err)
	}

	return e.buf.String()
}

func New(cfg *ucfg.Config, fields []*field.Field) (emitter.Emitter, error) {
	c := config{}
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	e := emt{
		fields: fields,
		buf:    bytes.NewBuffer(nil),
	}
	e.encoder = json.NewEncoder(e.buf)

	if c.Pretty {
		e.encoder.SetIndent("", "  ")
	}

	return &e, nil
}

func init() {
	emitter.Register(Name, New)
}
