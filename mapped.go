package fitert

import "iter"

// Mapped2 returns s mapped by m.
func Mapped2[T1, T2, O1, O2 any](s iter.Seq2[T1, T2], m func(T1, T2) (O1, O2)) iter.Seq2[O1, O2] {
	return func(yield func(O1, O2) bool) {
		for e1, e2 := range s {
			o1, o2 := m(e1, e2)
			c := yield(o1, o2)
			if !c {
				return
			}
		}
	}
}

// Mapped returns s mapped by m.
func Mapped[T, O any](s iter.Seq[T], m func(T) O) iter.Seq[O] {
	return Flattened(
		Mapped2(
			Splitted(s, func(t T) (T, any) { return t, nil }),
			func(t T, _ any) (O, any) { return m(t), nil },
		),
	)
}
