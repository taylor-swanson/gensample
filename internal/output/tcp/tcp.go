// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package tcp

import (
	"bufio"
	"fmt"
	"net"

	"github.com/elastic/go-ucfg"

	"github.com/taylor-swanson/gensample/internal/output"
)

const (
	Name = "tcp"
)

type out struct {
	Address      string `config:"address" validate:"required"`
	Network      string `config:"network"`
	OctetFraming bool   `config:"octet_framing"`
	Delimiter    string `config:"delimiter"`

	conn net.Conn
	buf  *bufio.Writer
}

func (o *out) Write(p []byte) (int, error) {
	var j int
	var k int
	var err error

	o.buf.Reset(o.conn)
	if o.OctetFraming {
		k, err = fmt.Fprintf(o.buf, "%d ", len(p))
		if err != nil {
			return k, err
		}
		j, err = o.buf.Write(p)
	} else {
		k, err = o.buf.Write(p)
		if err != nil {
			return j, err
		}
		k, err = fmt.Fprint(o.buf, o.Delimiter)
	}

	err = o.buf.Flush()

	return j + k, err
}

func (o *out) Close() error {
	return o.conn.Close()
}

func (o *out) NewInterval() error {
	return nil
}

func New(cfg *ucfg.Config) (output.Output, error) {
	o := out{
		Network:   "tcp",
		Delimiter: "\n",
	}
	if err := cfg.Unpack(&o); err != nil {
		return nil, err
	}

	var err error
	o.conn, err = net.Dial(o.Network, o.Address)
	if err != nil {
		return nil, fmt.Errorf("output.tcp: failed to connect: %w", err)
	}
	o.buf = bufio.NewWriter(o.conn)

	return &o, nil
}

func init() {
	output.Register(Name, New)
}
