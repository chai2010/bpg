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
	_ Image = (*RGB)(nil)
)

type RGB struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

func (p *RGB) Init(pix []uint8, stride int, rect image.Rectangle) *RGB {
	*p = RGB{
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

func (p *RGB) BaseType() image.Image { return p }
func (p *RGB) Pix() []byte           { return p.M.Pix }
func (p *RGB) Stride() int           { return p.M.Stride }
func (p *RGB) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB) Channels() int         { return 3 }
func (p *RGB) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGB) ColorModel() color.Model { return color.RGBAModel }

func (p *RGB) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB) At(x, y int) color.Color {
	c := p.RGBAt(x, y)
	return color.RGBA{
		R: c[0],
		G: c[1],
		B: c[2],
		A: 0xFF,
	}
}

func (p *RGB) RGBAt(x, y int) [3]uint8 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [3]uint8{}
	}
	i := p.PixOffset(x, y)
	return [3]uint8{
		p.M.Pix[i+0],
		p.M.Pix[i+1],
		p.M.Pix[i+2],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*3
}

func (p *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	p.M.Pix[i+0] = c1.R
	p.M.Pix[i+1] = c1.G
	p.M.Pix[i+2] = c1.B
	return
}

func (p *RGB) SetRGB(x, y int, c RGBColor) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = c.R
	p.M.Pix[i+1] = c.G
	p.M.Pix[i+2] = c.B
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB{
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
func (p *RGB) Opaque() bool {
	return true
}

// NewRGB returns a new RGB with the given bounds.
func NewRGB(r image.Rectangle) *RGB {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 3*w*h)
	return &RGB{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: 3 * w,
			Rect:   r,
		},
	}
}

func NewRGBFromImage(m image.Image) *RGB {
	if m, ok := m.(*RGB); ok {
		return m
	}

	// try `Image` interface
	if x, ok := m.(Image); ok {
		// try original type
		if m, ok := x.BaseType().(*RGB); ok {
			return m
		}
		// create new image with `x.Pix()`
		if x.Channels() == 3 && x.Depth() == reflect.Uint8 {
			return new(RGB).Init(x.Pix(), x.Stride(), x.Rect())
		}
	}

	// convert to RGB
	b := m.Bounds()
	rgb := NewRGB(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, _ := m.At(x, y).RGBA()
			rgb.SetRGB(x, y, RGBColor{
				uint8(pr >> 8),
				uint8(pg >> 8),
				uint8(pb >> 8),
			})
		}
	}
	return rgb
}
