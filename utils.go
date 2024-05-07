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

// MapKeys returns a slice of the keys in the map.
// Alternatively available from golang.org/x/exp/maps.
func MapKeys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues returns a slice of the values in the map.
// Alternatively available from golang.org/x/exp/maps.
func MapValues[T comparable, U any](m map[T]U) []U {
	values := make([]U, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
