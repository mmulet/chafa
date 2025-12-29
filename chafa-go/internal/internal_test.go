// Copyright (C) 2018-2025 Hans Petter Jansson
//
// This file is part of Chafa, a program that shows pictures on text terminals.

package internal

import (
	"strings"
	"testing"
)

func TestPackUnpackColor(t *testing.T) {
	original := Color{Ch: [4]uint8{255, 128, 64, 200}}
	
	packed := PackColor(&original)
	unpacked := UnpackColor(packed)
	
	if unpacked != original {
		t.Errorf("Pack/Unpack mismatch: got %v, expected %v", unpacked, original)
	}
}

func TestColorAverage2(t *testing.T) {
	colorA := Color{Ch: [4]uint8{100, 150, 200, 255}}
	colorB := Color{Ch: [4]uint8{200, 100, 50, 255}}
	
	avg := ColorAverage2(colorA, colorB)
	
	// Average should be approximately halfway between the two
	if avg.Ch[0] < 100 || avg.Ch[0] > 200 {
		t.Errorf("Average R channel out of range: %d", avg.Ch[0])
	}
}

func TestColorAccumAdd(t *testing.T) {
	accum := &ColorAccum{}
	color := &Color{Ch: [4]uint8{10, 20, 30, 40}}
	
	ColorAccumAdd(accum, color)
	
	if accum.Ch[0] != 10 || accum.Ch[1] != 20 || accum.Ch[2] != 30 || accum.Ch[3] != 40 {
		t.Errorf("ColorAccumAdd failed: got %v", accum.Ch)
	}
	
	ColorAccumAdd(accum, color)
	
	if accum.Ch[0] != 20 || accum.Ch[1] != 40 || accum.Ch[2] != 60 || accum.Ch[3] != 80 {
		t.Errorf("ColorAccumAdd second addition failed: got %v", accum.Ch)
	}
}

func TestColorAccumDivScalar(t *testing.T) {
	accum := &ColorAccum{Ch: [4]int16{100, 200, 300, 400}}
	
	ColorAccumDivScalar(accum, 10)
	
	if accum.Ch[0] != 10 || accum.Ch[1] != 20 || accum.Ch[2] != 30 || accum.Ch[3] != 40 {
		t.Errorf("ColorAccumDivScalar failed: got %v", accum.Ch)
	}
}

func TestColorDiffFast(t *testing.T) {
	colorA := &Color{Ch: [4]uint8{100, 100, 100, 255}}
	colorB := &Color{Ch: [4]uint8{110, 100, 100, 255}}
	
	diff := ColorDiffFast(colorA, colorB)
	
	// Difference should be 10^2 = 100
	if diff != 100 {
		t.Errorf("ColorDiffFast failed: got %d, expected 100", diff)
	}
	
	// Test with identical colors
	diff = ColorDiffFast(colorA, colorA)
	if diff != 0 {
		t.Errorf("ColorDiffFast with same color should be 0, got %d", diff)
	}
}

func TestColorRGBToDIN99D(t *testing.T) {
	// Test with red color
	red := &Color{Ch: [4]uint8{255, 0, 0, 255}}
	din99 := ColorRGBToDIN99D(red)
	
	// DIN99d should have different values than RGB
	if din99.Ch[0] == red.Ch[0] && din99.Ch[1] == red.Ch[1] && din99.Ch[2] == red.Ch[2] {
		t.Error("DIN99d conversion should change color values")
	}
	
	// Alpha should be preserved
	if din99.Ch[3] != red.Ch[3] {
		t.Errorf("Alpha channel not preserved: got %d, expected %d", din99.Ch[3], red.Ch[3])
	}
	
	// Test with white
	white := &Color{Ch: [4]uint8{255, 255, 255, 255}}
	din99White := ColorRGBToDIN99D(white)
	
	// White should have high L value
	if din99White.Ch[0] < 200 {
		t.Errorf("White should have high L value in DIN99d, got %d", din99White.Ch[0])
	}
	
	// Test with black
	black := &Color{Ch: [4]uint8{0, 0, 0, 255}}
	din99Black := ColorRGBToDIN99D(black)
	
	// Black should have low L value
	if din99Black.Ch[0] > 50 {
		t.Errorf("Black should have low L value in DIN99d, got %d", din99Black.Ch[0])
	}
}

func TestBase64Encode(t *testing.T) {
	b64 := NewBase64()
	var out strings.Builder
	
	// Test encoding a simple string
	input := []byte("Hello")
	b64.Encode(&out, input)
	b64.EncodeEnd(&out)
	
	result := out.String()
	
	// "Hello" in base64 should be "SGVsbG8="
	expected := "SGVsbG8="
	if result != expected {
		t.Errorf("Base64 encoding failed: got %q, expected %q", result, expected)
	}
}

func TestBase64EncodeChunked(t *testing.T) {
	b64 := NewBase64()
	var out strings.Builder
	
	// Test encoding in chunks
	b64.Encode(&out, []byte("He"))
	b64.Encode(&out, []byte("llo"))
	b64.EncodeEnd(&out)
	
	result := out.String()
	expected := "SGVsbG8="
	
	if result != expected {
		t.Errorf("Chunked base64 encoding failed: got %q, expected %q", result, expected)
	}
}

func TestBase64EncodeEmpty(t *testing.T) {
	b64 := NewBase64()
	var out strings.Builder
	
	b64.Encode(&out, []byte{})
	b64.EncodeEnd(&out)
	
	result := out.String()
	
	if result != "" {
		t.Errorf("Empty input should produce empty output, got %q", result)
	}
}

func TestBase64EncodeLong(t *testing.T) {
	b64 := NewBase64()
	var out strings.Builder
	
	// Test with longer data
	input := []byte("The quick brown fox jumps over the lazy dog")
	b64.Encode(&out, input)
	b64.EncodeEnd(&out)
	
	result := out.String()
	
	// Should produce valid base64
	if len(result) == 0 {
		t.Error("Long input should produce non-empty output")
	}
	
	// Base64 output should be longer than input and multiple of 4
	if len(result) < len(input) || len(result)%4 != 0 {
		t.Errorf("Invalid base64 output length: %d", len(result))
	}
}
