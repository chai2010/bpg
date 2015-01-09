// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bpg

/*
#cgo CFLAGS : -I./internal/libbpg-0.9.4
#cgo LDFLAGS: -L. -lbpg

#include <stdint.h>
#include <libbpg.h>
*/
import "C"
import (
	"unsafe"
)

type BPGDecoderContext C.BPGDecoderContext

type BPGImageFormatEnum C.BPGImageFormatEnum

const (
	BPG_FORMAT_GRAY      = BPGImageFormatEnum(C.BPG_FORMAT_GRAY)
	BPG_FORMAT_420       = BPGImageFormatEnum(C.BPG_FORMAT_420)
	BPG_FORMAT_422       = BPGImageFormatEnum(C.BPG_FORMAT_422)
	BPG_FORMAT_444       = BPGImageFormatEnum(C.BPG_FORMAT_444)
	BPG_FORMAT_420_VIDEO = BPGImageFormatEnum(C.BPG_FORMAT_420_VIDEO)
	BPG_FORMAT_422_VIDEO = BPGImageFormatEnum(C.BPG_FORMAT_422_VIDEO)
)

type BPGColorSpaceEnum C.BPGColorSpaceEnum

const (
	BPG_CS_YCbCr        = BPGColorSpaceEnum(C.BPG_CS_YCbCr)
	BPG_CS_RGB          = BPGColorSpaceEnum(C.BPG_CS_RGB)
	BPG_CS_YCgCo        = BPGColorSpaceEnum(C.BPG_CS_YCgCo)
	BPG_CS_YCbCr_BT709  = BPGColorSpaceEnum(C.BPG_CS_YCbCr_BT709)
	BPG_CS_YCbCr_BT2020 = BPGColorSpaceEnum(C.BPG_CS_YCbCr_BT2020)
	BPG_CS_COUNT        = BPGColorSpaceEnum(C.BPG_CS_COUNT)
)

// typedef struct {
//     int width;
//     int height;
//     int format; /* see BPGImageFormatEnum */
//     int has_alpha; /* TRUE if an alpha plane is present */
//     int color_space; /* see BPGColorSpaceEnum */
//     int bit_depth;
//     int premultiplied_alpha; /* TRUE if the color is alpha premultiplied */
//     int has_w_plane; /* TRUE if a W plane is present (for CMYK encoding) */
//     int limited_range; /* TRUE if limited range for the color */
// } BPGImageInfo;

type BPGImageInfo C.BPGImageInfo

type BPGExtensionTagEnum C.BPGExtensionTagEnum

const (
	BPG_EXTENSION_TAG_EXIF      = BPGExtensionTagEnum(C.BPG_EXTENSION_TAG_EXIF)
	BPG_EXTENSION_TAG_ICCP      = BPGExtensionTagEnum(C.BPG_EXTENSION_TAG_ICCP)
	BPG_EXTENSION_TAG_XMP       = BPGExtensionTagEnum(C.BPG_EXTENSION_TAG_XMP)
	BPG_EXTENSION_TAG_THUMBNAIL = BPGExtensionTagEnum(C.BPG_EXTENSION_TAG_THUMBNAIL)
)

// typedef struct BPGExtensionData {
//     BPGExtensionTagEnum tag;
//     uint32_t buf_len;
//     uint8_t *buf;
//     struct BPGExtensionData *next;
// } BPGExtensionData;

type BPGExtensionData C.BPGExtensionData

type BPGDecoderOutputFormat C.BPGDecoderOutputFormat

const (
	BPG_OUTPUT_FORMAT_RGB24  = BPGDecoderOutputFormat(C.BPG_OUTPUT_FORMAT_RGB24)
	BPG_OUTPUT_FORMAT_RGBA32 = BPGDecoderOutputFormat(C.BPG_OUTPUT_FORMAT_RGBA32)
	BPG_OUTPUT_FORMAT_RGB48  = BPGDecoderOutputFormat(C.BPG_OUTPUT_FORMAT_RGB48)
	BPG_OUTPUT_FORMAT_RGBA64 = BPGDecoderOutputFormat(C.BPG_OUTPUT_FORMAT_RGBA64)
)

const (
	BPG_DECODER_INFO_BUF_SIZE = int(C.BPG_DECODER_INFO_BUF_SIZE)
)

func bpg_decoder_open() *BPGDecoderContext {
	p := C.bpg_decoder_open()
	return (*BPGDecoderContext)(p)
}

/* If enable is true, extension data are kept during the image
   decoding and can be accessed after bpg_decoder_decode() with
   bpg_decoder_get_extension(). By default, the extension data are
   discarded. */
func bpg_decoder_keep_extension_data(s *BPGDecoderContext, enable int) {
	C.bpg_decoder_keep_extension_data(
		(*C.BPGDecoderContext)(s),
		(C.int)(enable),
	)
}

/* return 0 if 0K, < 0 if error */
func bpg_decoder_decode(s *BPGDecoderContext, buf unsafe.Pointer, buf_len int) int {
	rv := C.bpg_decoder_decode(
		(*C.BPGDecoderContext)(s),
		(*C.uint8_t)(buf),
		(C.int)(buf_len),
	)
	return int(rv)
}

/* Return the first element of the extension data list */
func bpg_decoder_get_extension_data(s *BPGDecoderContext) *BPGExtensionData {
	p := C.bpg_decoder_get_extension_data(
		(*C.BPGDecoderContext)(s),
	)
	return (*BPGExtensionData)(p)
}

/* return 0 if 0K, < 0 if error */
func bpg_decoder_get_info(s *BPGDecoderContext, p *BPGImageInfo) int {
	rv := C.bpg_decoder_get_info(
		(*C.BPGDecoderContext)(s),
		(*C.BPGImageInfo)(p),
	)
	return int(rv)
}

/* return 0 if 0K, < 0 if error */
func bpg_decoder_start(s *BPGDecoderContext, out_fmt BPGDecoderOutputFormat) int {
	rv := C.bpg_decoder_start(
		(*C.BPGDecoderContext)(s),
		(C.BPGDecoderOutputFormat)(out_fmt),
	)
	return int(rv)
}

/* return 0 if 0K, < 0 if error */
func bpg_decoder_get_line(s *BPGDecoderContext, buf unsafe.Pointer) int {
	rv := C.bpg_decoder_get_line(
		(*C.BPGDecoderContext)(s),
		buf,
	)
	return int(rv)
}

func bpg_decoder_close(s *BPGDecoderContext) {
	C.bpg_decoder_close(
		(*C.BPGDecoderContext)(s),
	)
}

/* only useful for low level access to the image data */
func bpg_decoder_get_data(s *BPGDecoderContext, plane int) (pline_size int, buf unsafe.Pointer) {
	panic("cgo disable pass go address to C!")
	var c_pline_size C.int
	rv := C.bpg_decoder_get_data(
		(*C.BPGDecoderContext)(s),
		(*C.int)(&c_pline_size), // cgo disable pass go address to C!!!!
		(C.int)(plane),
	)
	pline_size = int(c_pline_size)
	buf = unsafe.Pointer(rv)
	return
}

/* Get information from the start of the image data in 'buf' (at least
   min(BPG_DECODER_INFO_BUF_SIZE, file_size) bytes must be given).

   If pfirst_md != NULL, the extension data are also parsed and the
   first element of the list is returned in *pfirst_md. The list must
   be freed with bpg_decoder_free_extension_data().

   Return 0 if OK, < 0 if unrecognized data. */
func bpg_decoder_get_info_from_buf(
	p *BPGImageInfo,
	pfirst_md **BPGExtensionData,
	buf unsafe.Pointe,
	buf_len int,
) int {
	panic("TODO")
}

/* Free the extension data returned by bpg_decoder_get_info_from_buf() */
func bpg_decoder_free_extension_data(first_md *BPGExtensionData) {
	C.bpg_decoder_free_extension_data(
		(*C.BPGExtensionData)(first_md),
	)
}
