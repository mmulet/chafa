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
	"math"
)

// Square returns the square of a number.
func Square(n int) int {
	return n * n
}

// SquareFloat returns the square of a float.
func SquareFloat(n float64) float64 {
	return n * n
}

// alignDim aligns a dimension within a container.
func alignDim(align int, srcSize, destSize int) int {
	if srcSize > destSize {
		return 0
	}
	
	switch align {
	case 0: // AlignStart
		return 0
	case 1: // AlignEnd
		return destSize - srcSize
	case 2: // AlignCenter
		return (destSize - srcSize) / 2
	default:
		panic("invalid alignment")
	}
}

// TuckAndAlign calculates position and size for placing a source rectangle
// in a destination rectangle according to alignment and tucking policy.
func TuckAndAlign(srcWidth, srcHeight, destWidth, destHeight int,
	halign, valign int, tuck int) (ofsX, ofsY, width, height int) {
	
	switch tuck {
	case 0: // TuckStretch
		ofsX = 0
		ofsY = 0
		width = destWidth
		height = destHeight
		
	case 2: // TuckShrinkToFit
		if srcWidth <= destWidth && srcHeight <= destHeight {
			// Image fits entirely in dest. Do alignment only, no scaling.
			width = srcWidth
			height = srcHeight
			break
		}
		fallthrough
		
	case 1: // TuckFit
		ratioW := float64(destWidth) / float64(srcWidth)
		ratioH := float64(destHeight) / float64(srcHeight)
		
		minRatio := math.Min(ratioW, ratioH)
		
		width = int(math.Ceil(float64(srcWidth) * minRatio))
		height = int(math.Ceil(float64(srcHeight) * minRatio))
		
	default:
		panic("invalid tuck mode")
	}
	
	// Never exceed the dest size
	if width > destWidth {
		width = destWidth
	}
	if height > destHeight {
		height = destHeight
	}
	
	ofsX = alignDim(halign, width, destWidth)
	ofsY = alignDim(valign, height, destHeight)
	
	return
}

// RoundUpToMultipleOf rounds a value up to the nearest multiple of m.
func RoundUpToMultipleOf(value, m int) int {
	value = value + m - 1
	return value - (value % m)
}

// Min returns the minimum of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Clamp clamps a value between min and max.
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampFloat clamps a float value between min and max.
func ClampFloat(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
