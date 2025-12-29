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

// CanvasConfig holds the configuration for a Canvas.
type CanvasConfig struct {
	width                   int
	height                  int
	cellWidth               int
	cellHeight              int
	canvasMode              CanvasMode
	colorExtractor          ColorExtractor
	colorSpace              ColorSpace
	symbolMap               *SymbolMap
	fillSymbolMap           *SymbolMap
	transparencyThreshold   float32
	fgColor                 uint32
	bgColor                 uint32
	workFactor              float32
	preprocessingEnabled    bool
	ditherMode              DitherMode
	ditherGrainWidth        int
	ditherGrainHeight       int
	ditherIntensity         float32
	pixelMode               PixelMode
	optimizations           Optimizations
	fgOnlyEnabled           bool
	passthrough             Passthrough
}

// NewCanvasConfig creates a new CanvasConfig with default values.
func NewCanvasConfig() *CanvasConfig {
	return &CanvasConfig{
		width:                 80,
		height:                25,
		cellWidth:             8,
		cellHeight:            8,
		canvasMode:            CanvasModeTruecolor,
		colorExtractor:        ColorExtractorAverage,
		colorSpace:            ColorSpaceRGB,
		symbolMap:             NewSymbolMap(),
		fillSymbolMap:         nil,
		transparencyThreshold: 0.5,
		fgColor:               0xffffff,
		bgColor:               0x000000,
		workFactor:            1.0,
		preprocessingEnabled:  true,
		ditherMode:            DitherModeNone,
		ditherGrainWidth:      4,
		ditherGrainHeight:     4,
		ditherIntensity:       1.0,
		pixelMode:             PixelModeSymbols,
		optimizations:         OptimizationAll,
		fgOnlyEnabled:         false,
		passthrough:           PassthroughNone,
	}
}

// Copy creates a copy of the CanvasConfig.
func (cc *CanvasConfig) Copy() *CanvasConfig {
	if cc == nil {
		return nil
	}
	
	copy := *cc
	
	if cc.symbolMap != nil {
		copy.symbolMap = cc.symbolMap.Copy()
	}
	if cc.fillSymbolMap != nil {
		copy.fillSymbolMap = cc.fillSymbolMap.Copy()
	}
	
	return &copy
}

// GetGeometry returns the geometry of the canvas in character cells.
func (cc *CanvasConfig) GetGeometry() (width, height int) {
	return cc.width, cc.height
}

// SetGeometry sets the geometry of the canvas in character cells.
func (cc *CanvasConfig) SetGeometry(width, height int) {
	if width > 0 {
		cc.width = width
	}
	if height > 0 {
		cc.height = height
	}
}

// GetCellGeometry returns the geometry of each character cell in pixels.
//
// Since: 1.4
func (cc *CanvasConfig) GetCellGeometry() (cellWidth, cellHeight int) {
	return cc.cellWidth, cc.cellHeight
}

// SetCellGeometry sets the geometry of each character cell in pixels.
//
// Since: 1.4
func (cc *CanvasConfig) SetCellGeometry(cellWidth, cellHeight int) {
	if cellWidth > 0 {
		cc.cellWidth = cellWidth
	}
	if cellHeight > 0 {
		cc.cellHeight = cellHeight
	}
}

// GetCanvasMode returns the canvas mode.
func (cc *CanvasConfig) GetCanvasMode() CanvasMode {
	return cc.canvasMode
}

// SetCanvasMode sets the canvas mode.
func (cc *CanvasConfig) SetCanvasMode(mode CanvasMode) {
	cc.canvasMode = mode
}

// GetColorExtractor returns the color extractor method.
//
// Since: 1.4
func (cc *CanvasConfig) GetColorExtractor() ColorExtractor {
	return cc.colorExtractor
}

// SetColorExtractor sets the color extractor method.
//
// Since: 1.4
func (cc *CanvasConfig) SetColorExtractor(extractor ColorExtractor) {
	cc.colorExtractor = extractor
}

// GetColorSpace returns the color space used for color matching.
func (cc *CanvasConfig) GetColorSpace() ColorSpace {
	return cc.colorSpace
}

// SetColorSpace sets the color space used for color matching.
func (cc *CanvasConfig) SetColorSpace(colorSpace ColorSpace) {
	cc.colorSpace = colorSpace
}

// GetSymbolMap returns the symbol map.
func (cc *CanvasConfig) GetSymbolMap() *SymbolMap {
	return cc.symbolMap
}

// SetSymbolMap sets the symbol map.
func (cc *CanvasConfig) SetSymbolMap(symbolMap *SymbolMap) {
	if symbolMap != nil {
		cc.symbolMap = symbolMap.Copy()
	}
}

// GetFillSymbolMap returns the fill symbol map.
func (cc *CanvasConfig) GetFillSymbolMap() *SymbolMap {
	return cc.fillSymbolMap
}

// SetFillSymbolMap sets the fill symbol map.
func (cc *CanvasConfig) SetFillSymbolMap(fillSymbolMap *SymbolMap) {
	if fillSymbolMap != nil {
		cc.fillSymbolMap = fillSymbolMap.Copy()
	} else {
		cc.fillSymbolMap = nil
	}
}

// GetTransparencyThreshold returns the transparency threshold.
func (cc *CanvasConfig) GetTransparencyThreshold() float32 {
	return cc.transparencyThreshold
}

// SetTransparencyThreshold sets the transparency threshold.
func (cc *CanvasConfig) SetTransparencyThreshold(threshold float32) {
	cc.transparencyThreshold = threshold
}

// GetFGColor returns the foreground color as a packed RGB value.
func (cc *CanvasConfig) GetFGColor() uint32 {
	return cc.fgColor
}

// SetFGColor sets the foreground color as a packed RGB value.
func (cc *CanvasConfig) SetFGColor(color uint32) {
	cc.fgColor = color & 0xffffff
}

// GetBGColor returns the background color as a packed RGB value.
func (cc *CanvasConfig) GetBGColor() uint32 {
	return cc.bgColor
}

// SetBGColor sets the background color as a packed RGB value.
func (cc *CanvasConfig) SetBGColor(color uint32) {
	cc.bgColor = color & 0xffffff
}

// GetWorkFactor returns the work factor.
func (cc *CanvasConfig) GetWorkFactor() float32 {
	return cc.workFactor
}

// SetWorkFactor sets the work factor.
func (cc *CanvasConfig) SetWorkFactor(workFactor float32) {
	if workFactor > 0 {
		cc.workFactor = workFactor
	}
}

// GetPreprocessingEnabled returns whether preprocessing is enabled.
func (cc *CanvasConfig) GetPreprocessingEnabled() bool {
	return cc.preprocessingEnabled
}

// SetPreprocessingEnabled sets whether preprocessing is enabled.
func (cc *CanvasConfig) SetPreprocessingEnabled(enabled bool) {
	cc.preprocessingEnabled = enabled
}

// GetDitherMode returns the dither mode.
//
// Since: 1.2
func (cc *CanvasConfig) GetDitherMode() DitherMode {
	return cc.ditherMode
}

// SetDitherMode sets the dither mode.
//
// Since: 1.2
func (cc *CanvasConfig) SetDitherMode(mode DitherMode) {
	cc.ditherMode = mode
}

// GetDitherGrainSize returns the dither grain size.
//
// Since: 1.2
func (cc *CanvasConfig) GetDitherGrainSize() (width, height int) {
	return cc.ditherGrainWidth, cc.ditherGrainHeight
}

// SetDitherGrainSize sets the dither grain size.
//
// Since: 1.2
func (cc *CanvasConfig) SetDitherGrainSize(width, height int) {
	if width > 0 {
		cc.ditherGrainWidth = width
	}
	if height > 0 {
		cc.ditherGrainHeight = height
	}
}

// GetDitherIntensity returns the dither intensity.
//
// Since: 1.2
func (cc *CanvasConfig) GetDitherIntensity() float32 {
	return cc.ditherIntensity
}

// SetDitherIntensity sets the dither intensity.
//
// Since: 1.2
func (cc *CanvasConfig) SetDitherIntensity(intensity float32) {
	cc.ditherIntensity = intensity
}

// GetPixelMode returns the pixel mode.
//
// Since: 1.4
func (cc *CanvasConfig) GetPixelMode() PixelMode {
	return cc.pixelMode
}

// SetPixelMode sets the pixel mode.
//
// Since: 1.4
func (cc *CanvasConfig) SetPixelMode(mode PixelMode) {
	cc.pixelMode = mode
}

// GetOptimizations returns the optimizations flags.
//
// Since: 1.6
func (cc *CanvasConfig) GetOptimizations() Optimizations {
	return cc.optimizations
}

// SetOptimizations sets the optimizations flags.
//
// Since: 1.6
func (cc *CanvasConfig) SetOptimizations(optimizations Optimizations) {
	cc.optimizations = optimizations
}

// GetFGOnlyEnabled returns whether foreground-only mode is enabled.
//
// Since: 1.8
func (cc *CanvasConfig) GetFGOnlyEnabled() bool {
	return cc.fgOnlyEnabled
}

// SetFGOnlyEnabled sets whether foreground-only mode is enabled.
//
// Since: 1.8
func (cc *CanvasConfig) SetFGOnlyEnabled(enabled bool) {
	cc.fgOnlyEnabled = enabled
}

// GetPassthrough returns the passthrough mode.
//
// Since: 1.14
func (cc *CanvasConfig) GetPassthrough() Passthrough {
	return cc.passthrough
}

// SetPassthrough sets the passthrough mode.
//
// Since: 1.14
func (cc *CanvasConfig) SetPassthrough(passthrough Passthrough) {
	cc.passthrough = passthrough
}
