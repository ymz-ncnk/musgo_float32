package musgo_float32

import "testing"

func TestMusgoFloat32(t *testing.T) {
	var n float32 = 5.9
	buf := make([]byte, Float32(n).SizeMUS())
	Float32(n).MarshalMUS(buf)

	var an float32
	(*Float32)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
