package udp

import (
	"fmt"
	"net"
	"strconv"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/output"
)

const (
	Name = "udp"
)

type out struct {
	Host    string `config:"host" validate:"required"`
	Port    int    `config:"port" validate:"required, min=1, max=65535"`
	Network string `config:"network" validate:"required"`

	conn net.Conn
}

func (o *out) Write(p []byte) (int, error) {
	return o.conn.Write(p)
}

func (o *out) Close() error {
	return o.conn.Close()
}

func (o *out) NewInterval() error {
	return nil
}

func New(cfg *ucfg.Config) (output.Output, error) {
	o := out{
		Network: "udp",
	}
	if err := cfg.Unpack(&o); err != nil {
		return nil, err
	}

	var err error
	o.conn, err = net.Dial(o.Network, net.JoinHostPort(o.Host, strconv.Itoa(o.Port)))
	if err != nil {
		return nil, fmt.Errorf("output.udp: failed to connect: %w", err)
	}

	return &o, nil
}

func init() {
	output.Register(Name, New)
}
