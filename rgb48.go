// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bpg

import (
	"image"
	"image/color"
	"reflect"
)

var (
	_ Image = (*RGB48)(nil)
)

// RGB48 is an in-memory image whose At method returns color.RGB48 values.
type RGB48 struct {
	M struct {
		// Pix holds the image's pixels, in R, G, B order and big-endian format. The pixel at
		// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*6].
		Pix []byte
		// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
		Stride int
		// Rect is the image's bounds.
		Rect image.Rectangle
	}
}

func (p *RGB48) Init(pix []uint8, stride int, rect image.Rectangle) *RGB48 {
	*p = RGB48{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		},
	}
	return p
}

func (p *RGB48) BaseType() image.Image { return p }
func (p *RGB48) Pix() []byte           { return p.M.Pix }
func (p *RGB48) Stride() int           { return p.M.Stride }
func (p *RGB48) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB48) Channels() int         { return 3 }
func (p *RGB48) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *RGB48) ColorModel() color.Model { return RGB48Model }

func (p *RGB48) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB48) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return RGB48Color{}
	}
	i := p.PixOffset(x, y)
	return RGB48Color{
		uint16(p.M.Pix[i+0])<<8 | uint16(p.M.Pix[i+1]),
		uint16(p.M.Pix[i+2])<<8 | uint16(p.M.Pix[i+3]),
		uint16(p.M.Pix[i+4])<<8 | uint16(p.M.Pix[i+5]),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB48) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*6
}

func (p *RGB48) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := RGB48Model.Convert(c).(RGB48Color)
	p.M.Pix[i+0] = uint8(c1.R >> 8)
	p.M.Pix[i+1] = uint8(c1.R)
	p.M.Pix[i+2] = uint8(c1.G >> 8)
	p.M.Pix[i+3] = uint8(c1.G)
	p.M.Pix[i+4] = uint8(c1.B >> 8)
	p.M.Pix[i+5] = uint8(c1.B)
}

func (p *RGB48) SetRGB48(x, y int, c RGB48Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = uint8(c.R >> 8)
	p.M.Pix[i+1] = uint8(c.R)
	p.M.Pix[i+2] = uint8(c.G >> 8)
	p.M.Pix[i+3] = uint8(c.G)
	p.M.Pix[i+4] = uint8(c.B >> 8)
	p.M.Pix[i+5] = uint8(c.B)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB48{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB48{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    p.M.Pix[i:],
			Stride: p.M.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB48) Opaque() bool {
	return true
}

// NewRGB48 returns a new RGB48 with the given bounds.
func NewRGB48(r image.Rectangle) *RGB48 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 6*w*h)
	return new(RGB48).Init(pix, 6*w, r)
}

func NewRGB48FromImage(m image.Image) *RGB48 {
	if m, ok := m.(*RGB48); ok {
		return m
	}

	// try `Image` interface
	if x, ok := m.(Image); ok {
		// try original type
		if m, ok := x.BaseType().(*RGB48); ok {
			return m
		}
		// create new image with `x.Pix()`
		if x.Channels() == 3 && x.Depth() == reflect.Uint16 {
			return new(RGB48).Init(x.Pix(), x.Stride(), x.Rect())
		}
	}

	// convert to RGB48
	b := m.Bounds()
	rgb48 := NewRGB48(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, _ := m.At(x, y).RGBA()
			rgb48.SetRGB48(x, y, RGB48Color{
				uint16(pr),
				uint16(pg),
				uint16(pb),
			})
		}
	}
	return rgb48
}
