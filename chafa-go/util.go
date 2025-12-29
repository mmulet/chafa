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

import (
	"math"
)

// CalcCanvasGeometry calculates an optimal geometry for a Canvas given the width
// and height of an input image, maximum width and height of the canvas, font
// ratio, zoom and stretch preferences.
//
// srcWidth and srcHeight are the dimensions of the source image.
// destWidth and destHeight are pointers to the maximum dimensions of the canvas
// in character cells. These will be replaced by the calculated values. If one or
// both is negative, they will be calculated based on the remaining parameters
// and aspect ratio.
//
// fontRatio is the font's width divided by its height. 0.5 is a typical value.
// zoom, if true, will upscale the image to fit maximum dimensions.
// stretch, if true, will ignore the aspect ratio of the source.
func CalcCanvasGeometry(srcWidth, srcHeight int, destWidth, destHeight *int,
	fontRatio float64, zoom, stretch bool) {
	
	if srcWidth < 0 || srcHeight < 0 {
		panic("source dimensions must be non-negative")
	}
	
	if fontRatio <= 0.0 {
		panic("font ratio must be positive")
	}
	
	dw := -1
	dh := -1
	
	if destWidth != nil {
		dw = *destWidth
	}
	if destHeight != nil {
		dh = *destHeight
	}
	
	// If any dimension is explicitly set to zero, width and height will both be zero
	if srcWidth == 0 || srcHeight == 0 || dw == 0 || dh == 0 {
		if destWidth != nil {
			*destWidth = 0
		}
		if destHeight != nil {
			*destHeight = 0
		}
		return
	}
	
	// If both output dimensions are unspecified, make them 1/8 of their
	// corresponding input dimensions, rounding up and accounting for font ratio.
	// Both dimensions will be >= 1.
	if dw < 0 && dh < 0 {
		if destWidth != nil {
			*destWidth = (srcWidth + 7) / 8
			if *destWidth < 1 {
				*destWidth = 1
			}
		}
		
		if destHeight != nil {
			*destHeight = int(float64((srcHeight+7)/8)*fontRatio + 0.5)
			if *destHeight < 1 {
				*destHeight = 1
			}
		}
		
		return
	}
	
	if !zoom {
		if dw > srcWidth {
			dw = srcWidth
		}
		if dh > srcHeight {
			dh = srcHeight
		}
	}
	
	if !stretch || dw < 0 || dh < 0 {
		srcAspect := float64(srcWidth) / float64(srcHeight)
		destAspect := (float64(dw) / float64(dh)) * fontRatio
		
		if dw < 1 {
			dw = int(math.Ceil(float64(dh) * (srcAspect / fontRatio)))
		} else if dh < 1 {
			dh = int(math.Ceil((float64(dw) / srcAspect) * fontRatio))
		} else if srcAspect > destAspect {
			dh = int(math.Ceil(float64(dw) * (fontRatio / srcAspect)))
		} else {
			dw = int(math.Ceil(float64(dh) * (srcAspect / fontRatio)))
		}
	}
	
	// Clamp dest dimensions
	if dw < 1 {
		dw = 1
	}
	if dh < 1 {
		dh = 1
	}
	
	if destWidth != nil && *destWidth > 0 {
		if dw > *destWidth {
			dw = *destWidth
		}
	}
	if destHeight != nil && *destHeight > 0 {
		if dh > *destHeight {
			dh = *destHeight
		}
	}
	
	if destWidth != nil {
		*destWidth = dw
	}
	if destHeight != nil {
		*destHeight = dh
	}
}
