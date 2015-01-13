// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bpg

import (
	"image"
	"image/color"
	"io"
	"io/ioutil"
)

// DecodeConfig returns the color model and dimensions of a WEBP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	header := make([]byte, maxHeaderSize)
	n, err := r.Read(header)
	if err != nil && err != io.EOF {
		return
	}
	header, err = header[:n], nil
	info, err := DecodeInfo(header)
	if err != nil {
		return
	}
	config.Width = int(info.Width)
	config.Height = int(info.Height)
	if info.BitDepth == 8 {
		config.ColorModel = color.RGBAModel
	} else {
		config.ColorModel = color.RGBA64Model
	}
	return
}

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader) (m image.Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	info, err := DecodeInfo(data)
	if err != nil {
		return
	}
	if info.BitDepth == 8 {
		m, err = DecodeRGBA32(data)
		return
	} else {
		m, err = DecodeRGBA64(data)
		return
	}
}

func init() {
	image.RegisterFormat("bpg", headerMagic, Decode, DecodeConfig)
}
