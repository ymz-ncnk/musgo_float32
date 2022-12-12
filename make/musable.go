//go:build ignore

package main

import (
	"reflect"

	"github.com/ymz-ncnk/musgo"
	mgf "github.com/ymz-ncnk/musgo_float32"
)

func main() {
	musGo, err := musgo.New()
	if err != nil {
		panic(err)
	}
	unsafe := false
	err = musGo.Generate(reflect.TypeOf((*mgf.Float32)(nil)).Elem(), unsafe)
	if err != nil {
		panic(err)
	}
}
