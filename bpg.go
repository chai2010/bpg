// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bpg

import (
	"image"
)

func DecodeInfo(data []byte) (info FormatInfo, err error) {
	d := NewDecoder()
	defer d.Close()
	if err = d.Decode(data); err != nil {
		return
	}
	info, err = d.GetInfo()
	return
}

func DecodeExtension(data []byte) (ext []Extension, err error) {
	d := NewDecoder()
	defer d.Close()
	if err = d.Decode(data); err != nil {
		return
	}
	ext, err = d.GetExtension()
	return
}

func DecodeRGB24(data []byte) (m image.Image, err error) {
	d := NewDecoder()
	defer d.Close()
	if err = d.Decode(data); err != nil {
		return
	}
	m, err = d.GetImage(OutputFormatRGB24)
	return
}
func DecodeRGB48(data []byte) (m image.Image, err error) {
	d := NewDecoder()
	defer d.Close()
	if err = d.Decode(data); err != nil {
		return
	}
	m, err = d.GetImage(OutputFormatRGB48)
	return
}
func DecodeRGBA32(data []byte) (m image.Image, err error) {
	d := NewDecoder()
	defer d.Close()
	if err = d.Decode(data); err != nil {
		return
	}
	m, err = d.GetImage(OutputFormatRGBA32)
	return
}
func DecodeRGBA64(data []byte) (m image.Image, err error) {
	d := NewDecoder()
	defer d.Close()
	if err = d.Decode(data); err != nil {
		return
	}
	m, err = d.GetImage(OutputFormatRGBA64)
	return
}

type Decoder struct {
	*cgoBPGDecoderContext
}

func NewDecoder(keepExtionsin ...bool) *Decoder {
	pp := bpg_decoder_open()
	if len(keepExtionsin) > 0 {
		bpg_decoder_keep_extension_data(pp, keepExtionsin[0])
	}
	return &Decoder{pp}
}

func (p *Decoder) Close() {
	if p != nil && p.cgoBPGDecoderContext != nil {
		bpg_decoder_close(p.cgoBPGDecoderContext)
		p.cgoBPGDecoderContext = nil
	}
}

func (p *Decoder) Decode(data []byte) (err error) {
	return bpg_decoder_decode(p.cgoBPGDecoderContext, data)
}

func (p *Decoder) GetInfo() (info FormatInfo, err error) {
	return bpg_decoder_get_info(p.cgoBPGDecoderContext)
}

func (p *Decoder) GetExtension() (ext []Extension, err error) {
	return bpg_decoder_get_extension(p.cgoBPGDecoderContext)
}

func (p *Decoder) GetFrameDuration() (num, den int) {
	return bpg_decoder_get_frame_duration(p.cgoBPGDecoderContext)
}

func (p *Decoder) GetImage(outFormat OutputFormat) (m image.Image, err error) {
	return bpg_decoder_get_image(p.cgoBPGDecoderContext, outFormat)
}
