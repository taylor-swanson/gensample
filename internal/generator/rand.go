// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package generator

import (
	"encoding/binary"
	"math/rand/v2"
)

type RandReader struct {
	Rand *rand.Rand
}

func (r *RandReader) Read(p []byte) (n int, err error) {
	pos := 0
	for pos < len(p) {
		if pos+8 <= len(p) {
			rng := r.Rand.Uint64()
			binary.LittleEndian.PutUint64(p[pos:pos+8], rng)
			pos += 8
		} else if pos+4 <= len(p) {
			rng := r.Rand.Uint32()
			binary.LittleEndian.PutUint32(p[pos:pos+4], rng)
			pos += 4
		} else if pos+2 <= len(p) {
			rng := uint16(r.Rand.Uint32() % 0xFFFF)
			binary.LittleEndian.PutUint16(p[pos:pos+2], rng)
			pos += 2
		} else if pos+1 <= len(p) {
			rng := uint8(r.Rand.Uint32() % 0xFF)
			p[pos] = rng
			pos += 1
		}
	}

	return len(p), nil
}
