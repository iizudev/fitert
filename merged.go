package fitert

import "iter"

// Merged returns an iter.Seq from iter.Seq2 by merging
// the two value into a single one.
func Merged[T1, T2, R any](s iter.Seq2[T1, T2], m func(T1, T2) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for e1, e2 := range s {
			v := m(e1, e2)
			c := yield(v)
			if !c {
				break
			}
		}
	}
}
