// Copyright (C) 2018-2025 Hans Petter Jansson
//
// This file is part of Chafa, a program that shows pictures on text terminals.

// Package main provides a simple example of using the chafa-go library.
package main

import (
	"fmt"
	
	"github.com/mmulet/chafa/chafa-go"
)

func main() {
	fmt.Println("Chafa Go Library Example")
	fmt.Println("Version:", chafa.Version)
	fmt.Println()
	
	// Create a canvas configuration
	config := chafa.NewCanvasConfig()
	
	// Set some parameters
	config.SetGeometry(80, 25)
	config.SetCanvasMode(chafa.CanvasModeTruecolor)
	config.SetColorSpace(chafa.ColorSpaceRGB)
	
	w, h := config.GetGeometry()
	fmt.Printf("Canvas geometry: %d x %d\n", w, h)
	fmt.Printf("Canvas mode: %v\n", config.GetCanvasMode())
	fmt.Printf("Color space: %v\n", config.GetColorSpace())
	fmt.Println()
	
	// Create a canvas
	canvas := chafa.NewCanvas(config)
	if canvas != nil {
		fmt.Println("✓ Canvas created successfully")
	}
	
	// Create a symbol map
	symbolMap := chafa.NewSymbolMap()
	
	// Add some symbol tags
	symbolMap.AddByTags(chafa.SymbolTagBlock | chafa.SymbolTagBorder)
	fmt.Println("✓ Symbol map created with block and border symbols")
	
	fmt.Println()
	fmt.Println("Note: Full implementation is still in progress.")
	fmt.Println("This example demonstrates the basic API structure.")
}
