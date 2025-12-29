// Copyright (C) 2018-2025 Hans Petter Jansson
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
	"math"
)

// Palette index constants
const (
	PaletteIndexTransparent = 256
	PaletteIndexFG          = 257
	PaletteIndexBG          = 258
	PaletteIndexMax         = 259
)

// Color pair indexes (BG/FG indexes must be 0 and 1 respectively)
const (
	ColorPairBG = 0
	ColorPairFG = 1
)

// Color represents a color in a color space agnostic format.
// Contains 4 channels (R, G, B, A or transformed color space values).
type Color struct {
	Ch [4]uint8
}

// ColorPair represents a foreground and background color pair.
type ColorPair struct {
	Colors [2]Color
}

// ColorAccum is used for accumulating color values.
type ColorAccum struct {
	Ch [4]int16
}

// Pixel represents a single pixel with color information.
type Pixel struct {
	Col Color
}

// ColorCandidates represents color selection candidate pairs.
type ColorCandidates struct {
	Index [2]int16
	Error [2]int
}

// PaletteColor stores a color in multiple color spaces.
type PaletteColor struct {
	Col [2]Color // One for each color space (RGB, DIN99D)
}

// PackColor packs a color into a uint32.
func PackColor(color *Color) uint32 {
	// Assumes each channel 0 <= value <= 255
	return uint32(color.Ch[0])<<16 |
		uint32(color.Ch[1])<<8 |
		uint32(color.Ch[2]) |
		uint32(color.Ch[3])<<24 // Alpha
}

// UnpackColor unpacks a uint32 into a color.
func UnpackColor(packed uint32) Color {
	return Color{
		Ch: [4]uint8{
			uint8((packed >> 16) & 0xff),
			uint8((packed >> 8) & 0xff),
			uint8(packed & 0xff),
			uint8((packed >> 24) & 0xff), // Alpha
		},
	}
}

// ColorFromU32 creates a color from a uint32.
func ColorFromU32(u32 uint32) Color {
	return UnpackColor(u32)
}

// ColorToU32 converts a color to a uint32.
func ColorToU32(col Color) uint32 {
	return PackColor(&col)
}

// ColorAverage2 returns the average of two colors.
func ColorAverage2(colorA, colorB Color) Color {
	u32a := ColorToU32(colorA)
	u32b := ColorToU32(colorB)
	
	avg := ((u32a >> 1) & 0x7f7f7f7f) + ((u32b >> 1) & 0x7f7f7f7f)
	
	return ColorFromU32(avg)
}

// ColorAccumAdd adds a color to an accumulator.
func ColorAccumAdd(dest *ColorAccum, src *Color) {
	dest.Ch[0] += int16(src.Ch[0])
	dest.Ch[1] += int16(src.Ch[1])
	dest.Ch[2] += int16(src.Ch[2])
	dest.Ch[3] += int16(src.Ch[3])
}

// ColorAccumDivScalar divides a color accumulator by a scalar.
func ColorAccumDivScalar(accum *ColorAccum, scalar int) {
	accum.Ch[0] /= int16(scalar)
	accum.Ch[1] /= int16(scalar)
	accum.Ch[2] /= int16(scalar)
	accum.Ch[3] /= int16(scalar)
}

// ColorDiffFast computes a fast color difference (squared Euclidean distance).
func ColorDiffFast(colA, colB *Color) int {
	d0 := int(colB.Ch[0]) - int(colA.Ch[0])
	d1 := int(colB.Ch[1]) - int(colA.Ch[1])
	d2 := int(colB.Ch[2]) - int(colA.Ch[2])
	
	return d0*d0 + d1*d1 + d2*d2
}

// Color space conversion structures
type colorRGBf struct {
	c [3]float64
}

type colorXYZ struct {
	c [3]float64
}

type colorLab struct {
	c [3]float64
}

const (
	xyzEpsilon = 216.0 / 24389.0
	xyzKappa   = 24389.0 / 27.0
)

// invertRGBChannelCompand performs gamma correction inverse.
func invertRGBChannelCompand(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.044, 2.4)
}

// convertRGBToXYZ converts from RGB to XYZ color space.
func convertRGBToXYZ(rgb *Color) colorXYZ {
	var rgbf colorRGBf
	var xyz colorXYZ
	
	for i := 0; i < 3; i++ {
		rgbf.c[i] = float64(rgb.Ch[i]) / 255.0
		rgbf.c[i] = invertRGBChannelCompand(rgbf.c[i])
	}
	
	// sRGB to XYZ conversion matrix
	xyz.c[0] = 0.4124564*rgbf.c[0] + 0.3575761*rgbf.c[1] + 0.1804375*rgbf.c[2]
	xyz.c[1] = 0.2126729*rgbf.c[0] + 0.7151522*rgbf.c[1] + 0.0721750*rgbf.c[2]
	xyz.c[2] = 0.0193339*rgbf.c[0] + 0.1191920*rgbf.c[1] + 0.9503041*rgbf.c[2]
	
	return xyz
}

// labF is the CIE L*a*b* f function.
func labF(v float64) float64 {
	if v > xyzEpsilon {
		return math.Cbrt(v)
	}
	return (xyzKappa*v + 16.0) / 116.0
}

// convertXYZToLab converts from XYZ to L*a*b* color space.
func convertXYZToLab(xyz *colorXYZ) colorLab {
	// D65 white point
	wp := colorXYZ{c: [3]float64{0.95047, 1.0, 1.08883}}
	var xyz2 colorXYZ
	var lab colorLab
	
	for i := 0; i < 3; i++ {
		xyz2.c[i] = labF(xyz.c[i] / wp.c[i])
	}
	
	lab.c[0] = 116.0*xyz2.c[1] - 16.0
	lab.c[1] = 500.0 * (xyz2.c[0] - xyz2.c[1])
	lab.c[2] = 200.0 * (xyz2.c[1] - xyz2.c[2])
	
	return lab
}

// ColorRGBToDIN99D converts a color from RGB to DIN99d color space.
// DIN99d is a perceptually uniform color space designed for better color matching.
func ColorRGBToDIN99D(rgb *Color) Color {
	xyz := convertRGBToXYZ(rgb)
	
	// Apply tristimulus-space correction term
	xyz.c[0] = 1.12*xyz.c[0] - 0.12*xyz.c[2]
	
	// Convert to L*a*b*
	lab := convertXYZToLab(&xyz)
	adjL := 325.22 * math.Log(1.0+0.0036*lab.c[0])
	
	// Intermediate parameters
	ee := 0.6427876096865393*lab.c[1] + 0.766044443118978*lab.c[2]
	f := 1.14 * (0.6427876096865393*lab.c[2] - 0.766044443118978*lab.c[1])
	G := math.Sqrt(ee*ee + f*f)
	
	// Hue/chroma
	C := 22.5 * math.Log(1.0+0.06*G)
	
	h := math.Atan2(f, ee) + 0.8726646 // 50 degrees
	for h < 0.0 {
		h += 6.283185 // 360 degrees
	}
	for h > 6.283185 {
		h -= 6.283185 // 360 degrees
	}
	
	// The final values should be in the range [0..255]
	din99 := Color{
		Ch: [4]uint8{
			uint8(adjL * 2.5),
			uint8(C*math.Cos(h)*2.5 + 128.0),
			uint8(C*math.Sin(h)*2.5 + 128.0),
			rgb.Ch[3],
		},
	}
	
	return din99
}
