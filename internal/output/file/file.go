package udp

import (
	"os"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/output"
)

const (
	Name = "file"
)

type out struct {
	Filename  string `config:"format"`
	Directory string `config:"directory"`
	Pattern   string `config:"pattern"`
	Delimiter string `config:"delimiter"`

	file *os.File
}

func (o *out) Write(p []byte) (int, error) {
	j, err := o.file.Write(p)
	if err != nil {
		return j, err
	}
	k, err := o.file.Write([]byte(o.Delimiter))

	return j + k, err
}

func (o *out) Close() error {
	return o.file.Close()
}

func (o *out) NewInterval() error {
	if o.Directory == "" && o.Pattern == "" {
		return nil
	}
	if err := o.Close(); err != nil {
		return err
	}
	newFile, err := os.CreateTemp(o.Directory, o.Pattern)
	if err != nil {
		return err
	}
	o.file = newFile
	return nil
}

func New(cfg *ucfg.Config) (output.Output, error) {
	var err error

	o := out{}
	if err = cfg.Unpack(&o); err != nil {
		return nil, err
	}

	if o.Directory != "" && o.Pattern != "" {
		if o.file, err = os.CreateTemp(o.Directory, o.Pattern); err != nil {
			return nil, err
		}
	}
	if o.Filename != "" {
		if o.file, err = os.Create(o.Filename); err != nil {
			return nil, err
		}
	}

	return &o, nil
}

func init() {
	output.Register(Name, New)
}
