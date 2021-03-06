// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// cgoSafePtr/cgoFreePtr removed!
//
// Go1.3: Changes to the garbage collector
// http://golang.org/doc/go1.3#garbage_collector
//
// Go1.6:
// https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md

package bpg

import (
	"image"
)

/*
#cgo CFLAGS : -I./internal/libbpg-0.9.5 -DCONFIG_BPG_VERSION=0.9.5 -DUSE_VAR_BIT_DEPTH -DUSE_PRED -DHAVE_AV_CONFIG_H -Wno-pointer-sign -Wno-dangling-else -Wno-logical-op-parentheses -Wno-switch
//#cgo LDFLAGS: -L. -lbpg

#include <stdint.h>
#include <stdlib.h>
#include <libbpg.h>

// return 0 if 0K, < 0 if error
struct cgo_bpg_decoder_get_info_return {
	int retCode;
	BPGImageInfo info;
} cgo_bpg_decoder_get_info(BPGDecoderContext* p) {
	struct cgo_bpg_decoder_get_info_return t;
	t.retCode = bpg_decoder_get_info(p, &t.info);
	return t;
}

// return the frame delay for animations as a fraction (*pnum) / (*pden)
// in seconds. In case there is no animation, 0 / 1 is returned.
struct cgo_bpg_decoder_get_frame_duration_return {
	int num;
	int den;
} cgo_bpg_decoder_get_frame_duration(BPGDecoderContext* p) {
	struct cgo_bpg_decoder_get_frame_duration_return t;
	bpg_decoder_get_frame_duration(p, &t.num, &t.den);
	return t;
}

// return 0 if 0K, < 0 if error
struct cgo_bpg_decoder_get_image_return {
	int retCode;
	BPGImageInfo info;
	int pixelSize;
	char* ptr;
	int ptr_size;
} cgo_bpg_decoder_get_image(BPGDecoderContext* p, BPGDecoderOutputFormat format) {
	struct cgo_bpg_decoder_get_image_return t;
	int i;

	// get info
	t.retCode = bpg_decoder_get_info(p, &t.info);
	if(t.retCode < 0) {
		return t;
	}
	if(t.info.width <= 0 || t.info.height <= 0) {
		t.retCode = -1; // bad size
		return t;
	}

	// check format
	switch(format) {
	case BPG_OUTPUT_FORMAT_RGB24:
		t.pixelSize = 3;
		break;
	case BPG_OUTPUT_FORMAT_RGBA32:
		t.pixelSize = 4;
		break;
	case BPG_OUTPUT_FORMAT_RGB48:
		t.pixelSize = 6;
		break;
	case BPG_OUTPUT_FORMAT_RGBA64:
		t.pixelSize = 8;
		break;
	default:
		t.retCode = -1; // bad format
		return t;
	}

	// prepare for loop
	t.retCode = bpg_decoder_start(p, format);
	if(t.retCode < 0) {
		return t;
	}
	t.ptr_size = t.pixelSize*t.info.width*t.info.height;
	t.ptr = malloc(t.ptr_size);
	if(t.ptr == NULL) {
		t.retCode = -1;
		return t;
	}

	// loop
	for(i = 0; i < t.info.height; ++i) {
		void* curLine = t.ptr + t.pixelSize*t.info.width*i;
		t.retCode = bpg_decoder_get_line(p, curLine);
		if(t.retCode < 0) {
			free(t.ptr);
			return t;
		}
	}

	// OK
	return t;
}

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// ----------------------------------------------------------------------------
// types
// ----------------------------------------------------------------------------

type (
	cgoBPGDecoderContext C.BPGDecoderContext
)

// ----------------------------------------------------------------------------
// const
// ----------------------------------------------------------------------------

// format
const (
	cgoFormatGRAY     = Format(C.BPG_FORMAT_GRAY)
	cgoFormat420      = Format(C.BPG_FORMAT_420)
	cgoFormat422      = Format(C.BPG_FORMAT_422)
	cgoFormat444      = Format(C.BPG_FORMAT_444)
	cgoFormat420Video = Format(C.BPG_FORMAT_420_VIDEO)
	cgoFormat422Video = Format(C.BPG_FORMAT_422_VIDEO)
)

// color space
const (
	cgoColorSpaceYCbCr       = ColorSpace(C.BPG_CS_YCbCr)
	cgoColorSpaceRGB         = ColorSpace(C.BPG_CS_RGB)
	cgoColorSpaceYCgCo       = ColorSpace(C.BPG_CS_YCgCo)
	cgoColorSpaceYCbCrBT709  = ColorSpace(C.BPG_CS_YCbCr_BT709)
	cgoColorSpaceYCbCrBT2020 = ColorSpace(C.BPG_CS_YCbCr_BT2020)
	cgoColorSpaceMax         = ColorSpace(C.BPG_CS_COUNT)
)

const (
	cgoExtensionTagEXIF        = ExtensionTag(C.BPG_EXTENSION_TAG_EXIF)
	cgoExtensionTagICCP        = ExtensionTag(C.BPG_EXTENSION_TAG_ICCP)
	cgoExtensionTagXMP         = ExtensionTag(C.BPG_EXTENSION_TAG_XMP)
	cgoExtensionTagTHUMBNAIL   = ExtensionTag(C.BPG_EXTENSION_TAG_THUMBNAIL)
	cgoExtensionTagAnimControl = ExtensionTag(C.BPG_EXTENSION_TAG_ANIM_CONTROL)
)

const (
	cgoOutputFormatRGB24  = OutputFormat(C.BPG_OUTPUT_FORMAT_RGB24)
	cgoOutputFormatRGBA32 = OutputFormat(C.BPG_OUTPUT_FORMAT_RGBA32)
	cgoOutputFormatRGB48  = OutputFormat(C.BPG_OUTPUT_FORMAT_RGB48)
	cgoOutputFormatRGBA64 = OutputFormat(C.BPG_OUTPUT_FORMAT_RGBA64)
	cgoOutputFormatCMYK32 = OutputFormat(C.BPG_OUTPUT_FORMAT_CMYK32)
	cgoOutputFormatCMYK64 = OutputFormat(C.BPG_OUTPUT_FORMAT_CMYK64)
)

const (
	cgoDecoderInfoBufSize = int(C.BPG_DECODER_INFO_BUF_SIZE)
)

// ----------------------------------------------------------------------------
// func
// ----------------------------------------------------------------------------

// open/close
func bpg_decoder_open() *cgoBPGDecoderContext {
	p := C.bpg_decoder_open()
	return (*cgoBPGDecoderContext)(p)
}
func bpg_decoder_close(p *cgoBPGDecoderContext) {
	C.bpg_decoder_close(
		(*C.BPGDecoderContext)(p),
	)
}

// If enable is true, extension data are kept during the image
// decoding and can be accessed after bpg_decoder_decode() with
// bpg_decoder_get_extension(). By default, the extension data are
// discarded.
func bpg_decoder_keep_extension_data(s *cgoBPGDecoderContext, enabled bool) {
	if enabled {
		C.bpg_decoder_keep_extension_data(
			(*C.BPGDecoderContext)(s),
			1,
		)
	} else {
		C.bpg_decoder_keep_extension_data(
			(*C.BPGDecoderContext)(s),
			0,
		)
	}
}

// decode
func bpg_decoder_decode(p *cgoBPGDecoderContext, data []byte) (err error) {
	if len(data) == 0 {
		err = errors.New("bpg: bpg_decoder_decode: bad arguments")
		return
	}

	rv := C.bpg_decoder_decode(
		(*C.BPGDecoderContext)(p),
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		(C.int)(len(data)),
	)
	if rv < 0 {
		err = fmt.Errorf("bpg: bpg_decoder_decode, errcode = %d", rv)
		return
	}
	return
}

// get info
func bpg_decoder_get_info(p *cgoBPGDecoderContext) (info FormatInfo, err error) {
	rv := C.cgo_bpg_decoder_get_info(
		(*C.BPGDecoderContext)(p),
	)
	if rv.retCode < 0 {
		err = fmt.Errorf("bpg: bpg_decoder_get_info, errcode = %d", rv.retCode)
		return
	}
	info = FormatInfo{
		Width:              uint32(rv.info.width),
		Height:             uint32(rv.info.height),
		Format:             Format(rv.info.format),
		HasAlpha:           bool(rv.info.has_alpha != 0),
		ColorSpace:         ColorSpace(rv.info.color_space),
		BitDepth:           uint8(rv.info.bit_depth),
		PremultipliedAlpha: bool(rv.info.premultiplied_alpha != 0),
		HasWPlane:          bool(rv.info.has_w_plane != 0),
		LimitedRange:       bool(rv.info.limited_range != 0),
		HasAnimation:       bool(rv.info.has_animation != 0),
		LoopCount:          uint16(rv.info.loop_count),
	}
	return
}

// get extension
func bpg_decoder_get_extension(p *cgoBPGDecoderContext) (ext []Extension, err error) {
	first := C.bpg_decoder_get_extension_data(
		(*C.BPGDecoderContext)(p),
	)
	for x := first; x != nil; x = x.next {
		ext = append(ext, Extension{
			Tag:  ExtensionTag(x.tag),
			Data: C.GoBytes(unsafe.Pointer(x.buf), C.int(x.buf_len)),
		})
	}
	return
}

// return the frame delay for animations as a fraction (num) / (den)
// in seconds. In case there is no animation, 0 / 1 is returned.
func bpg_decoder_get_frame_duration(p *cgoBPGDecoderContext) (num, den int) {
	rv := C.cgo_bpg_decoder_get_frame_duration(
		(*C.BPGDecoderContext)(p),
	)
	num = int(rv.num)
	den = int(rv.den)
	return
}

// decode pixel
func bpg_decoder_get_image(p *cgoBPGDecoderContext, format OutputFormat) (m image.Image, err error) {
	rv := C.cgo_bpg_decoder_get_image(
		(*C.BPGDecoderContext)(p),
		(C.BPGDecoderOutputFormat)(format),
	)
	if rv.retCode < 0 {
		err = fmt.Errorf("bpg: bpg_decoder_get_image, errcode = %d", rv.retCode)
		return
	}

	pix := make([]byte, rv.ptr_size)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(pix):len(pix)])
	C.free(unsafe.Pointer(rv.ptr))

	rect := image.Rect(0, 0, int(rv.info.width), int(rv.info.height))
	stride := int(rv.info.width) * int(rv.pixelSize)

	switch format {
	case OutputFormatRGB24:
		m = new(RGB).Init(pix, stride, rect)
		return
	case OutputFormatRGB48:
		m = new(RGB48).Init(pix, stride, rect)
		return
	case OutputFormatRGBA32:
		m = &image.RGBA{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		}
		return
	case OutputFormatRGBA64:
		m = &image.RGBA64{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		}
		return
	case OutputFormatCMYK32:
		err = fmt.Errorf("bpg: bpg_decoder_get_image, unsupport OutputFormatCMYK32")
		return
	case OutputFormatCMYK64:
		err = fmt.Errorf("bpg: bpg_decoder_get_image, unsupport OutputFormatCMYK64")
		return
	default:
		panic("bpg: bpg_decoder_get_image, unreachable")
	}
}

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
