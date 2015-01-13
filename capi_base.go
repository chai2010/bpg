// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate stringer -type=Format       -output=z_gen_format_string.go
//go:generate stringer -type=ColorSpace   -output=z_gen_color_space_string.go
//go:generate stringer -type=ExtensionTag -output=z_gen_extension_tag_string.go
//go:generate stringer -type=OutputFormat -output=z_gen_output_format_string.go

package bpg

const (
	headerMagic = "BPG\xfb"
)

const (
	maxHeaderSize = 1024
)

type Format uint8

const (
	FormatGRAY     Format = iota
	Format420             // chroma at offset (0.5, 0.5) (JPEG)
	Format422             // chroma at offset (0.5, 0) (JPEG)
	Format444             //
	Format420Video        // chroma at offset (0, 0.5) (MPEG2)
	Format422Video        // chroma at offset (0, 0) (MPEG2)
	FormatMax
)

type ColorSpace uint8

const (
	ColorSpaceYCbCr ColorSpace = iota
	ColorSpaceRGB
	ColorSpaceYCgCo
	ColorSpaceYCbCrBT709
	ColorSpaceYCbCrBT2020
	ColorSpaceMax
)

type ExtensionTag uint8

const (
	_ ExtensionTag = iota
	ExtensionTagEXIF
	ExtensionTagICCP
	ExtensionTagXMP
	ExtensionTagTHUMBNAIL
	ExtensionTagAnimControl
)

type OutputFormat uint8

const (
	OutputFormatRGB24 OutputFormat = iota
	OutputFormatRGBA32
	OutputFormatRGB48
	OutputFormatRGBA64
	OutputFormatCMYK32
	OutputFormatCMYK64
)

type FormatInfo struct {
	Width              uint32
	Height             uint32
	Format             Format
	HasAlpha           bool // true if an alpha plane is present
	ColorSpace         ColorSpace
	BitDepth           uint8
	PremultipliedAlpha bool   // true if the color is alpha premultiplied
	HasWPlane          bool   // true if a W plane is present (for CMYK encoding)
	LimitedRange       bool   // true if limited range for the color
	HasAnimation       bool   // true if the image contains animations
	LoopCount          uint16 // animations: number of loop, 0 = infinity
}

type Extension struct {
	Tag  ExtensionTag
	Data []byte
}

type EncodeImage struct {
	Width              int
	Height             int
	Format             Format // x_VIDEO values are forbidden here
	CHHase             uint8  // 4:2:2 or 4:2:0 : give the horizontal chroma position. 0=MPEG2, 1=JPEG.
	HasAlpha           bool
	HasWPlane          bool
	LimitedRange       uint8
	PremultipliedAlpha bool
	ColorSpace         ColorSpace
	BitDepth           uint8
	PixelShift         uint8 // (1 << pixel_shift) bytes per pixel
	Data               [4][]byte
	LineSize           [4]int
}

type EncodeParams struct {
	Width                 int
	Height                int
	ChromaFormat          int  // 0-3
	BitDepth              int  // 8-14
	IntraOnly             int  // 0-1
	Quality               int  // quantizer 0-51
	Lossless              bool // 0-1 lossless mode
	SeiDecodedPictureHash int  // 0=no hash, 1=MD5 hash
	CompressLevel         int  // 1-9
	Verbose               bool
}
