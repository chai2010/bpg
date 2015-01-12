// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bpg

import (
	"image/color"
)

// RGB represents a traditional 24-bit fully opaque color,
// having 8 bits for each of red, green and blue.
type RGBColor struct {
	R, G, B uint8
}

func (c RGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xFFFF
	return
}

// RGB48 represents a 48-bit fully opaque color,
// having 16 bits for each of red, green and blue.
type RGB48Color struct {
	R, G, B uint16
}

func (c RGB48Color) RGBA() (r, g, b, a uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), 0xFFFF
}

var (
	RGBModel   color.Model = color.ModelFunc(rgbModel)
	RGB48Model color.Model = color.ModelFunc(rgb48Model)
)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGBColor); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGBColor{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

func rgb48Model(c color.Color) color.Color {
	if _, ok := c.(RGB48Color); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB48Color{uint16(r), uint16(g), uint16(b)}
}
