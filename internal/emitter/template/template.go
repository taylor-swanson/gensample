package template

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/elastic/go-ucfg"

	"gensample/internal/context"
	"gensample/internal/emitter"
	"gensample/internal/field"
)

const (
	Name = "template"
)

type config struct {
	Templates []string `config:"templates" validate:"required"`
}

type tmpl struct {
	tmpls  []*template.Template
	fields []*field.Field
}

func (t *tmpl) Emit(ctx *context.Context) string {
	values := make(map[string]string, len(t.fields))
	for _, v := range t.fields {
		values[v.Name] = v.Generator.Generate(ctx)
	}

	buf := bytes.NewBuffer(nil)

	err := t.tmpls[ctx.Current%len(t.tmpls)].Execute(buf, values)
	if err != nil {
		log.Printf("emitter.template: error: %v", err)
		return ""
	}

	return buf.String()
}

func New(cfg *ucfg.Config, fields []*field.Field) (emitter.Emitter, error) {
	c := config{}
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	t := tmpl{}
	for i, v := range c.Templates {
		tmplv, err := template.New("").Parse(v)
		if err != nil {
			return nil, fmt.Errorf("emitter.template: parse template %d error: %w", i, err)
		}
		t.tmpls = append(t.tmpls, tmplv)
	}
	t.fields = fields

	return &t, nil
}

func init() {
	emitter.Register(Name, New)
}
