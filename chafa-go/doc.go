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

/*
Package chafa is a terminal graphics library that converts images to text-based
representations using various terminal protocols (Sixel, Kitty, iTerm2) and
Unicode/ASCII characters.

This is a Go port of the original C library. The conversion aims to be as
literal as possible while maintaining idiomatic Go code.

Basic Usage

Create a canvas configuration:

	config := chafa.NewCanvasConfig()
	config.SetGeometry(80, 25)
	config.SetCanvasMode(chafa.CanvasModeTruecolor)

Create a canvas and render pixels:

	canvas := chafa.NewCanvas(config)
	canvas.DrawAllPixels(chafa.PixelRGBA8Premultiplied, pixels, width, height, rowstride)

Generate output:

	termInfo := chafa.NewTermInfo()
	output := canvas.Print(termInfo)
	fmt.Print(output)

Status

This Go port is currently under active development. Core types and structures
are in place, but many implementations are still pending.

License

LGPL 3.0 or later, same as the original Chafa library.
*/
package chafa
