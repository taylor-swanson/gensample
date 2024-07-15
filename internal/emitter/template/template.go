package template

import (
	"bytes"
	"gensample/internal/context"
	"log"
	"text/template"

	"github.com/elastic/go-ucfg"

	"gensample/internal/emitter"
	"gensample/internal/field"
)

const (
	Name = "template"
)

type config struct {
	Template string         `config:"template" validate:"required"`
	Fields   []*ucfg.Config `config:"fields" validate:"required"`
}

type tmpl struct {
	tmpl   *template.Template
	fields []*field.Field
}

func (t *tmpl) Emit(ctx *context.Context) string {
	values := make(map[string]string, len(t.fields))
	for _, v := range t.fields {
		values[v.Name] = v.Generator.Generate(ctx)
	}

	buf := bytes.NewBuffer(nil)

	err := t.tmpl.Execute(buf, values)
	if err != nil {
		log.Printf("emitter.template: error: %v", err)
		return ""
	}

	return buf.String()
}

func New(cfg *ucfg.Config) (emitter.Emitter, error) {
	c := config{}
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}

	var err error
	t := tmpl{}
	t.tmpl, err = template.New("").Parse(c.Template)
	if err != nil {
		return nil, err
	}

	for _, v := range c.Fields {
		f, err := field.New(v)
		if err != nil {
			return nil, err
		}
		t.fields = append(t.fields, f)
	}

	return &t, nil
}

func init() {
	emitter.Register(Name, New)
}
