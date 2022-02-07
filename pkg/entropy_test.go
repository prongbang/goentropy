package goentropy_test

import (
	"fmt"
	"testing"

	goentropy "github.com/prongbang/goentropy/pkg"
)

var e goentropy.Entropy

func init() {
	e = goentropy.New()
}

func TestEntropy(t *testing.T) {
	r := e.Calc("123456")
	if r > 2.5 {
		t.Errorf("Expected 2.5, got %f", r)
	}
}

func BenchmarkEntropy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = e.Calc(fmt.Sprintf("%d", i))
	}
}
