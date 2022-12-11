# MusGo_float32
Provides serialization of the Golang's `float32` type.

# How to use
Simply cast `float32` to `musgo_float32.Float32`:
```go
	var n float32 = 5.5
	// Marshal
	buf := make([]byte, musgo_float32.Float32(n).SizeMUS())
	musgo_float32.Float32(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_float32.Float32)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

