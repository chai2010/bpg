// generated by stringer -type=ColorSpace -output=z_gen_color_space_string.go; DO NOT EDIT

package bpg

import "fmt"

const _ColorSpace_name = "ColorSpaceYCbCrColorSpaceRGBColorSpaceYCgCoColorSpaceYCbCrBT709ColorSpaceYCbCrBT2020ColorSpaceMax"

var _ColorSpace_index = [...]uint8{15, 28, 43, 63, 84, 97}

func (i ColorSpace) String() string {
	if i >= ColorSpace(len(_ColorSpace_index)) {
		return fmt.Sprintf("ColorSpace(%d)", i)
	}
	hi := _ColorSpace_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _ColorSpace_index[i-1]
	}
	return _ColorSpace_name[lo:hi]
}
