package lynn

import (
	"math/rand/v2"
)

func Map[T, V any](s []T, fn func(T) V) []V {
	rs := make([]V, len(s))
	for i, e := range s {
		rs[i] = fn(e)
	}
	return rs
}

func Rand(min, max int) int {
	return rand.IntN(max-min) + min
}
