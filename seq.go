package fitert

import (
	"iter"
	"slices"
)

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func Seq2[T1, T2 any](e ...Pair[T1, T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		for _, p := range e {
			c := yield(p.First, p.Second)
			if !c {
				break
			}
		}
	}
}

func Seq[T any](e ...T) iter.Seq[T] {
	return Flattened(
		Seq2(
			slices.Collect(
				Mapped(
					slices.Values(e),
					func(e T) Pair[T, any] { return Pair[T, any]{e, nil} },
				),
			)...,
		),
	)
}
