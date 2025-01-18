package fitert

import "iter"

// Splitted returns single iter.Seq of T splitted into
// an iter.Seq2 of O1 and O2 using m.
func Splitted[T, O1, O2 any](s iter.Seq[T], m func(T) (O1, O2)) iter.Seq2[O1, O2] {
	return func(yield func(O1, O2) bool) {
		for e := range s {
			v1, v2 := m(e)
			c := yield(v1, v2)
			if !c {
				break
			}
		}
	}
}
