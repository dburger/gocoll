package gocoll

// MultiSet provides a map that keeps counts of elements.
type MultiSet[T comparable] struct {
	counts map[T]int
}

// NewMultiSet constructs and returns a MultiSet of the given type.
func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{make(map[T]int)}
}

// Add adds an element to the multiset and returns the new count.
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
// Remove removes the value from the MultiSet and returns the new count.
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
// Count returns the count of value in the MultiSet.
func (ms *MultiSet[T]) Count(val T) int {
	count, ok := ms.counts[val]
	if !ok {
		count = 0
	}
	return count
}
