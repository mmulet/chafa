// Copyright (C) 2018-2025 Hans Petter Jansson
//
// This file is part of Chafa, a program that shows pictures on text terminals.

package chafa

import (
	"testing"
)

func TestNewCanvasConfig(t *testing.T) {
	config := NewCanvasConfig()
	if config == nil {
		t.Fatal("NewCanvasConfig returned nil")
	}
	
	// Test default values
	if w, h := config.GetGeometry(); w != 80 || h != 25 {
		t.Errorf("Expected default geometry 80x25, got %dx%d", w, h)
	}
	
	if mode := config.GetCanvasMode(); mode != CanvasModeTruecolor {
		t.Errorf("Expected default canvas mode truecolor, got %v", mode)
	}
}

func TestCanvasConfigSettersGetters(t *testing.T) {
	config := NewCanvasConfig()
	
	// Test geometry
	config.SetGeometry(100, 50)
	if w, h := config.GetGeometry(); w != 100 || h != 50 {
		t.Errorf("Expected geometry 100x50, got %dx%d", w, h)
	}
	
	// Test canvas mode
	config.SetCanvasMode(CanvasModeIndexed256)
	if mode := config.GetCanvasMode(); mode != CanvasModeIndexed256 {
		t.Errorf("Expected canvas mode indexed256, got %v", mode)
	}
	
	// Test colors
	config.SetFGColor(0xff0000)
	if fg := config.GetFGColor(); fg != 0xff0000 {
		t.Errorf("Expected FG color 0xff0000, got 0x%x", fg)
	}
	
	config.SetBGColor(0x0000ff)
	if bg := config.GetBGColor(); bg != 0x0000ff {
		t.Errorf("Expected BG color 0x0000ff, got 0x%x", bg)
	}
}

func TestCanvasConfigCopy(t *testing.T) {
	config1 := NewCanvasConfig()
	config1.SetGeometry(100, 50)
	config1.SetCanvasMode(CanvasModeIndexed256)
	
	config2 := config1.Copy()
	if config2 == nil {
		t.Fatal("Copy returned nil")
	}
	
	if w, h := config2.GetGeometry(); w != 100 || h != 50 {
		t.Errorf("Expected copied geometry 100x50, got %dx%d", w, h)
	}
	
	if mode := config2.GetCanvasMode(); mode != CanvasModeIndexed256 {
		t.Errorf("Expected copied canvas mode indexed256, got %v", mode)
	}
	
	// Modify config1 and ensure config2 is independent
	config1.SetGeometry(200, 100)
	if w, h := config2.GetGeometry(); w != 100 || h != 50 {
		t.Errorf("Copy is not independent, expected 100x50, got %dx%d", w, h)
	}
}

func TestNewCanvas(t *testing.T) {
	config := NewCanvasConfig()
	canvas := NewCanvas(config)
	if canvas == nil {
		t.Fatal("NewCanvas returned nil")
	}
	
	canvasConfig := canvas.GetConfig()
	if canvasConfig == nil {
		t.Fatal("Canvas config is nil")
	}
}

func TestSymbolMap(t *testing.T) {
	sm := NewSymbolMap()
	if sm == nil {
		t.Fatal("NewSymbolMap returned nil")
	}
	
	// Test default builtin glyphs
	if !sm.GetAllowBuiltinGlyphs() {
		t.Error("Expected builtin glyphs to be allowed by default")
	}
	
	// Test setting builtin glyphs
	sm.SetAllowBuiltinGlyphs(false)
	if sm.GetAllowBuiltinGlyphs() {
		t.Error("Expected builtin glyphs to be disabled")
	}
}

func TestSymbolTags(t *testing.T) {
	// Test symbol tag constants
	if SymbolTagNone != 0 {
		t.Error("SymbolTagNone should be 0")
	}
	
	// Test combined tags
	if SymbolTagHalf != (SymbolTagHHalf | SymbolTagVHalf) {
		t.Error("SymbolTagHalf should be combination of HHalf and VHalf")
	}
	
	if SymbolTagAlnum != (SymbolTagAlpha | SymbolTagDigit) {
		t.Error("SymbolTagAlnum should be combination of Alpha and Digit")
	}
}

func TestPixelTypes(t *testing.T) {
	// Just verify the constants are defined
	types := []PixelType{
		PixelRGBA8Premultiplied,
		PixelBGRA8Premultiplied,
		PixelARGB8Premultiplied,
		PixelABGR8Premultiplied,
		PixelRGBA8Unassociated,
		PixelBGRA8Unassociated,
		PixelARGB8Unassociated,
		PixelABGR8Unassociated,
		PixelRGB8,
		PixelBGR8,
	}
	
	if len(types) != 10 {
		t.Errorf("Expected 10 pixel types, got %d", len(types))
	}
}

func TestOptimizationFlags(t *testing.T) {
	// Test that optimization flags can be combined
	opts := OptimizationReuseAttributes | OptimizationRepeatCells
	
	if opts&OptimizationReuseAttributes == 0 {
		t.Error("Expected OptimizationReuseAttributes to be set")
	}
	
	if opts&OptimizationRepeatCells == 0 {
		t.Error("Expected OptimizationRepeatCells to be set")
	}
	
	if opts&OptimizationSkipCells != 0 {
		t.Error("Expected OptimizationSkipCells to not be set")
	}
}
