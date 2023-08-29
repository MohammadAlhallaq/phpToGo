package generics

type Collection[T int | float64] []T

func (c Collection[T]) Sum() T {

	var t T

	for i := 0; i < len(c); i++ {
		t += c[i]
	}

	return t
}
