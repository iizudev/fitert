package fitert

import (
	"iter"
)

// First2 returns the first two elements of the iter.Seq2 s
// and a bool indicating whether the elements are present (the
// seq2 may be empty — in this case false is returned alogside
// zeroed T1 and T2).
func First2[T1, T2 any](s iter.Seq2[T1, T2]) (T1, T2, bool) {
	next, stop := iter.Pull2(s)
	defer stop()
	return next()
}

// First returns the first element of the iter.Seq s and a bool
// indicating whether the element is present (the seq may be empty —
// in this case false is returned alongside a zeroed T).
func First[T any](s iter.Seq[T]) (v T, ok bool) {
	v, _, ok = First2(Splitted(s, func(t T) (T, any) { return t, nil }))
	return
}
