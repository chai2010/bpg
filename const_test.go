// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bpg

import (
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		got, expect Format
	}{
		{FormatGRAY, cgoFormatGRAY},
		{Format420, cgoFormat420},
		{Format422, cgoFormat422},
		{Format444, cgoFormat444},
		{Format420Video, cgoFormat420Video},
		{Format422Video, cgoFormat422Video},
	}
	for i, v := range tests {
		if v.got != v.expect {
			t.Fatalf("%d: expect = %v, got = %v", i, v.got, v.expect)
		}
	}
}

func TestColorSpace(t *testing.T) {
	tests := []struct {
		got, expect ColorSpace
	}{
		{ColorSpaceYCbCr, cgoColorSpaceYCbCr},
		{ColorSpaceRGB, cgoColorSpaceRGB},
		{ColorSpaceYCgCo, cgoColorSpaceYCgCo},
		{ColorSpaceYCbCrBT709, cgoColorSpaceYCbCrBT709},
		{ColorSpaceYCbCrBT2020, cgoColorSpaceYCbCrBT2020},
		{ColorSpaceMax, cgoColorSpaceMax},
	}
	for i, v := range tests {
		if v.got != v.expect {
			t.Fatalf("%d: expect = %v, got = %v", i, v.got, v.expect)
		}
	}
}

func TestExtensionTag(t *testing.T) {
	tests := []struct {
		got, expect ExtensionTag
	}{
		{ExtensionTagEXIF, cgoExtensionTagEXIF},
		{ExtensionTagICCP, cgoExtensionTagICCP},
		{ExtensionTagXMP, cgoExtensionTagXMP},
		{ExtensionTagTHUMBNAIL, cgoExtensionTagTHUMBNAIL},
	}
	for i, v := range tests {
		if v.got != v.expect {
			t.Fatalf("%d: expect = %v, got = %v", i, v.got, v.expect)
		}
	}
}

func TestOutputFormat(t *testing.T) {
	tests := []struct {
		got, expect OutputFormat
	}{
		{OutputFormatRGB24, cgoOutputFormatRGB24},
		{OutputFormatRGBA32, cgoOutputFormatRGBA32},
		{OutputFormatRGB48, cgoOutputFormatRGB48},
		{OutputFormatRGBA64, cgoOutputFormatRGBA64},
	}
	for i, v := range tests {
		if v.got != v.expect {
			t.Fatalf("%d: expect = %v, got = %v", i, v.got, v.expect)
		}
	}
}
