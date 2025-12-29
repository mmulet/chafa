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

const (
	// SymbolWidthPixels is the width of a symbol in pixels.
	SymbolWidthPixels = 8
	// SymbolHeightPixels is the height of a symbol in pixels.
	SymbolHeightPixels = 8
)

// SymbolTags represents categories of symbols that can be selected.
type SymbolTags uint32

const (
	SymbolTagNone SymbolTags = 0

	SymbolTagSpace     SymbolTags = 1 << 0
	SymbolTagSolid     SymbolTags = 1 << 1
	SymbolTagStipple   SymbolTags = 1 << 2
	SymbolTagBlock     SymbolTags = 1 << 3
	SymbolTagBorder    SymbolTags = 1 << 4
	SymbolTagDiagonal  SymbolTags = 1 << 5
	SymbolTagDot       SymbolTags = 1 << 6
	SymbolTagQuad      SymbolTags = 1 << 7
	SymbolTagHHalf     SymbolTags = 1 << 8
	SymbolTagVHalf     SymbolTags = 1 << 9
	SymbolTagHalf      SymbolTags = SymbolTagHHalf | SymbolTagVHalf
	SymbolTagInverted  SymbolTags = 1 << 10
	SymbolTagBraille   SymbolTags = 1 << 11
	SymbolTagTechnical SymbolTags = 1 << 12
	SymbolTagGeometric SymbolTags = 1 << 13
	SymbolTagASCII     SymbolTags = 1 << 14
	SymbolTagAlpha     SymbolTags = 1 << 15
	SymbolTagDigit     SymbolTags = 1 << 16
	SymbolTagAlnum     SymbolTags = SymbolTagAlpha | SymbolTagDigit
	SymbolTagNarrow    SymbolTags = 1 << 17
	SymbolTagWide      SymbolTags = 1 << 18
	SymbolTagAmbiguous SymbolTags = 1 << 19
	SymbolTagUgly      SymbolTags = 1 << 20
	SymbolTagLegacy    SymbolTags = 1 << 21
	SymbolTagSextant   SymbolTags = 1 << 22
	SymbolTagWedge     SymbolTags = 1 << 23
	SymbolTagLatin     SymbolTags = 1 << 24
	SymbolTagImported  SymbolTags = 1 << 25
	SymbolTagOctant    SymbolTags = 1 << 26
	SymbolTagExtra     SymbolTags = 1 << 30
	SymbolTagBad       SymbolTags = SymbolTagAmbiguous | SymbolTagUgly
	SymbolTagAll       SymbolTags = ^(SymbolTagExtra | SymbolTagBad)
)

// SymbolMap represents a mapping of symbols to use for rendering.
type SymbolMap struct {
	// TODO: Internal structure to be implemented
	allowBuiltinGlyphs bool
	// Additional fields will be added during full conversion
}

// NewSymbolMap creates a new SymbolMap.
func NewSymbolMap() *SymbolMap {
	return &SymbolMap{
		allowBuiltinGlyphs: true,
	}
}

// Copy creates a copy of the SymbolMap.
func (sm *SymbolMap) Copy() *SymbolMap {
	if sm == nil {
		return nil
	}
	copy := *sm
	return &copy
}

// AddByTags adds symbols with the specified tags to the symbol map.
func (sm *SymbolMap) AddByTags(tags SymbolTags) {
	// TODO: Implementation
}

// RemoveByTags removes symbols with the specified tags from the symbol map.
func (sm *SymbolMap) RemoveByTags(tags SymbolTags) {
	// TODO: Implementation
}

// AddByRange adds symbols in the specified Unicode range to the symbol map.
//
// Since: 1.4
func (sm *SymbolMap) AddByRange(first, last rune) {
	// TODO: Implementation
}

// RemoveByRange removes symbols in the specified Unicode range from the symbol map.
//
// Since: 1.4
func (sm *SymbolMap) RemoveByRange(first, last rune) {
	// TODO: Implementation
}

// ApplySelectors applies a selector string to the symbol map.
// Returns an error if the selector string is invalid.
func (sm *SymbolMap) ApplySelectors(selectors string) error {
	// TODO: Implementation
	return nil
}

// GetAllowBuiltinGlyphs returns whether builtin glyphs are allowed.
//
// Since: 1.4
func (sm *SymbolMap) GetAllowBuiltinGlyphs() bool {
	return sm.allowBuiltinGlyphs
}

// SetAllowBuiltinGlyphs sets whether builtin glyphs are allowed.
//
// Since: 1.4
func (sm *SymbolMap) SetAllowBuiltinGlyphs(allow bool) {
	sm.allowBuiltinGlyphs = allow
}

// AddGlyph adds a custom glyph to the symbol map.
//
// Since: 1.4
func (sm *SymbolMap) AddGlyph(codePoint rune, pixelFormat PixelType, pixels []byte, width, height, rowstride int) {
	// TODO: Implementation
}

// GetGlyph retrieves a glyph from the symbol map.
// Returns the pixel data, dimensions, and whether the glyph was found.
//
// Since: 1.12
func (sm *SymbolMap) GetGlyph(codePoint rune, pixelFormat PixelType) (pixels []byte, width, height, rowstride int, ok bool) {
	// TODO: Implementation
	return nil, 0, 0, 0, false
}
