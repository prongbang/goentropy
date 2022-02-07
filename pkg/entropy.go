package goentropy

import (
	"math"
)

type Entropy interface {
	Calc(data string) float64
}

type entropy struct {
}

func (e *entropy) Calc(data string) float64 {

	mn := map[rune]float64{}
	for _, r := range data {
		mn[r]++
	}

	var hm float64
	for _, n := range mn {
		hm += n * math.Log2(n)
	}

	l := float64(len(data))
	r := math.Log2(l) - hm/l

	return r
}

func New() Entropy {
	return &entropy{}
}
