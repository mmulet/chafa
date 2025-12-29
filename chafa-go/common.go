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

package chafa

// PixelType represents pixel formats supported by Canvas and SymbolMap.
//
// Since: 1.4
type PixelType int

const (
	// 32 bits per pixel

	// PixelRGBA8Premultiplied is premultiplied RGBA, 8 bits per channel.
	PixelRGBA8Premultiplied PixelType = iota
	// PixelBGRA8Premultiplied is premultiplied BGRA, 8 bits per channel.
	PixelBGRA8Premultiplied
	// PixelARGB8Premultiplied is premultiplied ARGB, 8 bits per channel.
	PixelARGB8Premultiplied
	// PixelABGR8Premultiplied is premultiplied ABGR, 8 bits per channel.
	PixelABGR8Premultiplied

	// PixelRGBA8Unassociated is unassociated RGBA, 8 bits per channel.
	PixelRGBA8Unassociated
	// PixelBGRA8Unassociated is unassociated BGRA, 8 bits per channel.
	PixelBGRA8Unassociated
	// PixelARGB8Unassociated is unassociated ARGB, 8 bits per channel.
	PixelARGB8Unassociated
	// PixelABGR8Unassociated is unassociated ABGR, 8 bits per channel.
	PixelABGR8Unassociated

	// 24 bits per pixel

	// PixelRGB8 is packed RGB (no alpha), 8 bits per channel.
	PixelRGB8
	// PixelBGR8 is packed BGR (no alpha), 8 bits per channel.
	PixelBGR8

	// PixelMax is the last supported pixel type, plus one.
	PixelMax
)

// Align represents alignment options when placing an element within an area.
//
// Since: 1.14
type Align int

const (
	// AlignStart aligns flush with beginning of the area (top or left in LTR locales).
	AlignStart Align = iota
	// AlignEnd aligns flush with end of the area (bottom or right in LTR locales).
	AlignEnd
	// AlignCenter aligns in the middle of the area.
	AlignCenter

	// AlignMax is the last supported alignment, plus one.
	AlignMax
)

// Tuck represents resizing options when placing an element within an area.
// Usually used in conjunction with Align to control the padding.
//
// Since: 1.14
type Tuck int

const (
	// TuckStretch resizes element to fit the area exactly, changing its aspect ratio.
	TuckStretch Tuck = iota
	// TuckFit resizes element to fit the area, preserving its aspect ratio by adding padding.
	TuckFit
	// TuckShrinkToFit is like TuckFit, but prohibits enlargement.
	TuckShrinkToFit

	// TuckMax is the last supported tucking policy, plus one.
	TuckMax
)

// ColorExtractor specifies the method to use for extracting colors.
type ColorExtractor int

const (
	// ColorExtractorAverage uses the average colors of each symbol's coverage area.
	ColorExtractorAverage ColorExtractor = iota
	// ColorExtractorMedian uses the median colors of each symbol's coverage area.
	ColorExtractorMedian

	// ColorExtractorMax is the last supported color extractor plus one.
	ColorExtractorMax
)

// ColorSpace represents color spaces used for color matching.
type ColorSpace int

const (
	// ColorSpaceRGB is the RGB color space. Fast but imprecise.
	ColorSpaceRGB ColorSpace = iota
	// ColorSpaceDIN99D is the DIN99d color space. Slower, but good perceptual color precision.
	ColorSpaceDIN99D

	// ColorSpaceMax is the last supported color space plus one.
	ColorSpaceMax
)

// DitherMode represents dithering modes.
type DitherMode int

const (
	// DitherModeNone means no dithering.
	DitherModeNone DitherMode = iota
	// DitherModeOrdered is ordered dithering (Bayer or similar).
	DitherModeOrdered
	// DitherModeDiffusion is error diffusion dithering (Floyd-Steinberg or similar).
	DitherModeDiffusion
	// DitherModeNoise is noise pattern dithering (blue noise or similar).
	DitherModeNoise

	// DitherModeMax is the last supported dither mode plus one.
	DitherModeMax
)

// Optimizations flags for sequence optimization. When enabled, these may produce
// more compact output at the cost of reduced compatibility and increased CPU use.
// Output quality is unaffected.
type Optimizations int

const (
	// OptimizationReuseAttributes suppresses redundant SGR control sequences.
	OptimizationReuseAttributes Optimizations = 1 << 0
	// OptimizationSkipCells is reserved for future use.
	OptimizationSkipCells Optimizations = 1 << 1
	// OptimizationRepeatCells uses REP sequence to compress repeated runs of similar cells.
	OptimizationRepeatCells Optimizations = 1 << 2

	// OptimizationNone means all optimizations disabled.
	OptimizationNone Optimizations = 0
	// OptimizationAll means all optimizations enabled.
	OptimizationAll Optimizations = 0x7fffffff
)

// CanvasMode represents canvas color modes.
type CanvasMode int

const (
	// CanvasModeTruecolor is truecolor mode.
	CanvasModeTruecolor CanvasMode = iota
	// CanvasModeIndexed256 uses 256 colors.
	CanvasModeIndexed256
	// CanvasModeIndexed240 uses 256 colors, but avoids using the lower 16 whose values vary between terminal environments.
	CanvasModeIndexed240
	// CanvasModeIndexed16 uses 16 colors using the aixterm ANSI extension.
	CanvasModeIndexed16
	// CanvasModeFGBGBGFG uses default foreground and background colors, plus inversion.
	CanvasModeFGBGBGFG
	// CanvasModeFGBG uses default foreground and background colors. No ANSI codes will be used.
	CanvasModeFGBG
	// CanvasModeIndexed8 uses 8 colors, compatible with original ANSI X3.64.
	CanvasModeIndexed8
	// CanvasModeIndexed16_8 uses 16 FG colors (8 of which enabled with bold/bright) and 8 BG colors.
	CanvasModeIndexed16_8

	// CanvasModeMax is the last supported canvas mode plus one.
	CanvasModeMax
)

// PixelMode represents pixel encoding modes.
type PixelMode int

const (
	// PixelModeSymbols approximates pixel data using character symbols ("ANSI art").
	PixelModeSymbols PixelMode = iota
	// PixelModeSixels encodes pixel data as sixels.
	PixelModeSixels
	// PixelModeKitty encodes pixel data using the Kitty terminal protocol.
	PixelModeKitty
	// PixelModeITerm2 encodes pixel data using the iTerm2 terminal protocol.
	PixelModeITerm2

	// PixelModeMax is the last supported pixel mode plus one.
	PixelModeMax
)

// Passthrough represents passthrough modes for terminal multiplexers.
type Passthrough int

const (
	// PassthroughNone means no passthrough guards will be used.
	PassthroughNone Passthrough = iota
	// PassthroughScreen uses passthrough guards for GNU Screen.
	PassthroughScreen
	// PassthroughTmux uses passthrough guards for tmux.
	PassthroughTmux

	// PassthroughMax is the last supported passthrough mode plus one.
	PassthroughMax
)
