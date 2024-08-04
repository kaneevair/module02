package main

import (
	"fmt"
	"testing"
)

func BenchmarkWithPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := `a`
		fmt.Print(a)
		b := IntsCopyWithPointer(&a)
		*b += "g"
	}
}

func BenchmarkWithoutPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := `a`
		fmt.Print(a)
		b := IntsCopyWithoutPointer(a)
		b += "g"
	}
}

func IntsCopyWithPointer(src *string) *string {
	*src += "f"

	return src

}

func IntsCopyWithoutPointer(src string) string {
	src += "f"

	return src

}
