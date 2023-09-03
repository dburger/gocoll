package gocoll

type MultiSet[T comparable] struct {
	counts map[T]int
}

func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{make(map[T]int)}
}

func (ms *MultiSet[T]) Add(val T) int {
	count, ok := ms.counts[val]
	if !ok {
		count = 1
	} else {
		count++
	}
	ms.counts[val] = count
	return count
}

// TODO(dburger): test.
func (ms *MultiSet[T]) Remove(val T) int {
	count, ok := ms.counts[val]
	if !ok || count == 1 {
		delete(ms.counts, val)
	} else {
		count--
		ms.counts[val] = count
	}
	return count
}

// TODO(dburger): test.
func (ms *MultiSet[T]) Count(val T) int {
	count, ok := ms.counts[val]
	if !ok {
		count = 0
	}
	return count
}
