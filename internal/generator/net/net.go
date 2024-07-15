package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"

	"github.com/elastic/go-ucfg"
	"lukechampine.com/uint128"

	"gensample/internal/context"
	"gensample/internal/generator"
)

const (
	NamePort = "port"
	NameIP   = "ip"
	NameMAC  = "mac"
)

type port struct {
	Min int `config:"min" validate:"required, min=0, max=65535"`
	Max int `config:"max" validate:"required, min=0, max=65535"`
}

func (g *port) Validate() error {
	if g.Min > g.Max {
		return fmt.Errorf("port: invalid range, min > max (%d > %d)", g.Min, g.Max)
	}
	return nil
}

func (g *port) Generate(ctx *context.Context) string {
	return strconv.Itoa(ctx.Rand.IntN(g.Max-g.Min+1) + g.Min)
}

func NewPortGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := port{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	return &g, nil
}

type ip struct {
	CIDR string `config:"cidr" validate:"required"`

	cidr *net.IPNet
	ip6  bool
}

func (g *ip) Validate() error {
	_, _, err := net.ParseCIDR(g.CIDR)
	if err != nil {
		return fmt.Errorf("ip: unable to parse cidr: %s", err)
	}

	return nil
}

func (g *ip) Generate(ctx *context.Context) string {
	if g.ip6 {
		lower := ipToUint128(g.cidr.IP).Add64(1)
		upper := ipToUint128(g.cidr.Mask).Xor(uint128.Uint128{Lo: 0xFFFFFFFFFFFFFFFF, Hi: 0xFFFFFFFFFFFFFFFF}).Or(lower).Sub64(1)

		rng := uint128.New(ctx.Rand.Uint64(), ctx.Rand.Uint64())
		rng = rng.Mod(upper.Sub(lower)).Add(lower)

		v := uint128ToIP(rng)

		return v.String()
	}

	low := ipToUint32(g.cidr.IP) + 1
	mask := ipToUint32(g.cidr.Mask)
	high := ((mask ^ 0xFFFFFFFF) | low) - 1

	rng := ctx.Rand.Uint32()%(high-low) + low
	v := uint32ToIP(rng)

	return v.String()
}

func NewIPGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := ip{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	_, g.cidr, _ = net.ParseCIDR(g.CIDR)
	g.ip6 = len(g.cidr.IP) == net.IPv6len

	return &g, nil
}

type mac struct {
	Separator string `config:"separator"`
	Lowercase bool   `config:"lowercase"`
}

func (g *mac) Generate(ctx *context.Context) string {
	rng := ctx.Rand.Uint64()

	var str string
	for i := 0; i < 6; i++ {
		v := (rng >> (i * 8)) & 0xFF

		if g.Lowercase {
			str += fmt.Sprintf("%02x", v)
		} else {
			str += fmt.Sprintf("%02X", v)
		}
		if i < 5 {
			str += g.Separator
		}
	}

	return str
}

func NewMACGenerator(cfg *ucfg.Config) (generator.Generator, error) {
	g := mac{}
	if err := cfg.Unpack(&g); err != nil {
		return nil, err
	}

	if g.Separator == "" {
		g.Separator = "-"
	}

	return &g, nil
}

func ipToUint32(b []byte) uint32 {
	if len(b) == net.IPv6len {
		return binary.BigEndian.Uint32(b[12:16])
	}

	return binary.BigEndian.Uint32(b)
}

func uint32ToIP(n uint32) net.IP {
	ip := make([]byte, net.IPv4len)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

func ipToUint128(ip []byte) uint128.Uint128 {
	if len(ip) != net.IPv6len {
		return uint128.Uint128{}
	}

	return uint128.Uint128{
		Lo: binary.BigEndian.Uint64(ip[8:]),
		Hi: binary.BigEndian.Uint64(ip[:8]),
	}
}

func uint128ToIP(v uint128.Uint128) net.IP {
	ip := make([]byte, net.IPv6len)

	binary.BigEndian.PutUint64(ip[8:], v.Lo)
	binary.BigEndian.PutUint64(ip[:8], v.Hi)

	return ip
}

func init() {
	generator.Register(NamePort, NewPortGenerator)
	generator.Register(NameIP, NewIPGenerator)
	generator.Register(NameMAC, NewMACGenerator)
}
