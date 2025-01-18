package fitert

import (
	"iter"
)

// Find2 returns the first two elements of s that match the p predicate.
func Find2[T1, T2 any](s iter.Seq2[T1, T2], p func(*T1, *T2) bool) (T1, T2, bool) {
	return First2(Filtered2(s, p))
}

// Find returns the first element of s that matches the p predicate.
func Find[T any](s iter.Seq[T], p func(*T) bool) (v T, ok bool) {
	// explicit usage of Find2 because this function should only depend
	// on the Find2 implementation.
	v, _, ok = Find2(
		Splitted(s, func(t T) (T, any) { return t, nil }),
		func(t *T, _ *any) bool { return p(t) },
	)
	return
}
