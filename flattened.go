package fitert

import "iter"

// Flattened returns a flattened version of s, which only
// produces the first element of the iter.Seq2.
//
// This function can be combined with Swapped to achieve
// the same result but for the second element of the original
// seq2.
func Flattened[T1, T2 any](s iter.Seq2[T1, T2]) iter.Seq[T1] {
	return Merged(s, func(t T1, _ T2) T1 { return t })
}
