package stdout

import (
	"fmt"
	"os"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/output"
)

const (
	Name = "stdout"
)

type out struct {
}

func (o *out) Write(p []byte) (n int, err error) {
	return fmt.Fprintln(os.Stdout, string(p))
}

func (o *out) Close() error {
	return nil
}

func (o *out) NewInterval() error {
	return nil
}

func New(_ *ucfg.Config) (output.Output, error) {
	return &out{}, nil
}

func init() {
	output.Register(Name, New)
}
