package gocoll

// Supposedly slices.Repeat coming in 1.23.

// RepeatFunc returns a slice of the given length with elements computed by a function.
func RepeatFunc[T any](f func() T, n int) []T {
	var res = make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = f()
	}
	return res
}

// RepeatValues returns a slice of the given length with elements of the given value.
func RepeatValues[T any](val T, n int) []T {
	var res = make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = val
	}
	return res
}
