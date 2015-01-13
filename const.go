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
