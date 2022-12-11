// Code generated by musgen. DO NOT EDIT.

package musgo_float32

import (
	"math"

	"github.com/ymz-ncnk/musgo/errs"
)

// MarshalMUS fills buf with the MUS encoding of v.
func (v Float32) MarshalMUS(buf []byte) int {
	i := 0
	{
		uv := math.Float32bits(float32(v))
		uv = (uv << 16) | (uv >> 16)
		uv = ((uv << 8) & 0xFF00FF00) | ((uv >> 8) & 0x00FF00FF)
		{
			for uv >= 0x80 {
				buf[i] = byte(uv) | 0x80
				uv >>= 7
				i++
			}
			buf[i] = byte(uv)
			i++
		}
	}
	return i
}

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *Float32) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var uv uint32
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			}
			shift := 0
			done := false
			for l, b := range buf[i:] {
				if l == 4 && b > 15 {
					return i, errs.ErrOverflow
				}
				if b < 0x80 {
					uv = uv | uint32(b)<<shift
					done = true
					i += l + 1
					break
				}
				uv = uv | uint32(b&0x7F)<<shift
				shift += 7
			}
			if !done {
				return i, errs.ErrSmallBuf
			}
		}
		uv = (uv << 16) | (uv >> 16)
		uv = ((uv << 8) & 0xFF00FF00) | ((uv >> 8) & 0x00FF00FF)
		(*v) = Float32(math.Float32frombits(uv))
	}
	return i, err
}

// SizeMUS returns the size of the MUS-encoded v.
func (v Float32) SizeMUS() int {
	size := 0
	{
		uv := math.Float32bits(float32(v))
		uv = (uv << 16) | (uv >> 16)
		uv = ((uv << 8) & 0xFF00FF00) | ((uv >> 8) & 0x00FF00FF)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}

	}
	return size
}
