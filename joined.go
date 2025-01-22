package fitert

import (
	"iter"
	"slices"
)

// Joined2 returns a single iter.Seq2[T1, T2] of multiple provided sequences.
func Joined2[T1, T2 any](c ...iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		for _, i := range c {
			for e1, e2 := range i {
				c := yield(e1, e2)
				if !c {
					break
				}
			}
		}
	}
}

// Joined returns a single iter.Seq[T] of multiple provided sequences.
func Joined[T any](c ...iter.Seq[T]) iter.Seq[T] {
	return Flattened(
		Joined2(
			slices.Collect(
				Mapped(
					slices.Values(c),
					func(i iter.Seq[T]) iter.Seq2[T, any] {
						return Splitted(
							i,
							func(t T) (T, any) { return t, nil })
					},
				),
			)...,
		),
	)
}
