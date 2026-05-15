// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package udp

import (
	"fmt"
	"net"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/output"
)

const (
	Name = "udp"
)

type out struct {
	Address string `config:"address" validate:"required"`
	Network string `config:"network"`

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
	o.conn, err = net.Dial(o.Network, o.Address)
	if err != nil {
		return nil, fmt.Errorf("output.udp: failed to connect: %w", err)
	}

	return &o, nil
}

func init() {
	output.Register(Name, New)
}
