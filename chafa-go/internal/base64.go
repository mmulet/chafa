// Copyright (C) 2021-2025 Hans Petter Jansson
//
// This file is part of Chafa, a program that shows pictures on text terminals.
//
// Chafa is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Chafa is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Chafa.  If not, see <http://www.gnu.org/licenses/>.

package internal

import (
	"encoding/base64"
	"strings"
)

// Base64 handles base64 encoding with buffering for streaming.
type Base64 struct {
	// We turn 3-byte groups into 4-character base64 groups, so we
	// may need to buffer up to 2 bytes between batches
	buf    [2]byte
	bufLen int
}

// NewBase64 creates a new Base64 encoder.
func NewBase64() *Base64 {
	return &Base64{}
}

// encode3Bytes encodes 3 bytes to base64.
func encode3Bytes(bytes uint32) string {
	const base64Dict = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	
	return string([]byte{
		base64Dict[(bytes>>(3*6))&0x3f],
		base64Dict[(bytes>>(2*6))&0x3f],
		base64Dict[(bytes>>(1*6))&0x3f],
		base64Dict[bytes&0x3f],
	})
}

// Encode encodes data to base64, appending to the output builder.
func (b *Base64) Encode(out *strings.Builder, in []byte) {
	inLen := len(in)
	if inLen == 0 {
		return
	}
	
	inPos := 0
	
	if b.bufLen+inLen < 3 {
		copy(b.buf[b.bufLen:], in)
		b.bufLen += inLen
		return
	}
	
	var r uint32
	
	if b.bufLen == 1 {
		r = (uint32(b.buf[0]) << 16) | (uint32(in[0]) << 8) | uint32(in[1])
		inPos += 2
		out.WriteString(encode3Bytes(r))
	} else if b.bufLen == 2 {
		r = (uint32(b.buf[0]) << 16) | (uint32(b.buf[1]) << 8) | uint32(in[0])
		inPos++
		out.WriteString(encode3Bytes(r))
	}
	
	b.bufLen = 0
	
	for inLen-inPos >= 3 {
		r = (uint32(in[inPos]) << 16) | (uint32(in[inPos+1]) << 8) | uint32(in[inPos+2])
		out.WriteString(encode3Bytes(r))
		inPos += 3
	}
	
	for inPos < inLen {
		b.buf[b.bufLen] = in[inPos]
		b.bufLen++
		inPos++
	}
}

// EncodeEnd finalizes the base64 encoding, encoding any remaining buffered bytes.
func (b *Base64) EncodeEnd(out *strings.Builder) {
	const base64Dict = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	
	if b.bufLen == 1 {
		out.WriteByte(base64Dict[b.buf[0]>>2])
		out.WriteByte(base64Dict[(b.buf[0]<<4)&0x30])
		out.WriteString("==")
	} else if b.bufLen == 2 {
		out.WriteByte(base64Dict[b.buf[0]>>2])
		out.WriteByte(base64Dict[((b.buf[0]<<4)|(b.buf[1]>>4))&0x3f])
		out.WriteByte(base64Dict[(b.buf[1]<<2)&0x3c])
		out.WriteByte('=')
	}
	
	b.bufLen = 0
}

// EncodeStandard is a convenience function to encode data using Go's standard base64 encoder.
func EncodeStandard(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
