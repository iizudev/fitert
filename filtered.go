package fitert

import "iter"

// Filtered2 returns s only with the elements that pass
// the p predicate checks.
func Filtered2[T1, T2 any](s iter.Seq2[T1, T2], p func(*T1, *T2) bool) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		for e1, e2 := range s {
			if !p(&e1, &e2) {
				continue
			}
			c := yield(e1, e2)
			if !c {
				break
			}
		}
	}
}

// Filtered returns s only with the elements that pass
// the p predicate checks.
func Filtered[T any](s iter.Seq[T], p func(*T) bool) iter.Seq[T] {
	return Flattened(
		Filtered2(
			Splitted(s, func(t T) (T, any) { return t, nil }),
			func(t *T, _ *any) bool { return p(t) },
		),
	)
}
