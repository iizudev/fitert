package fitert

import "iter"

// Swapped returns s but with swapped elements.
func Swapped[T1, T2 any](s iter.Seq2[T1, T2]) iter.Seq2[T2, T1] {
	return Mapped2(s, func(t1 T1, t2 T2) (T2, T1) { return t2, t1 })
}
