package gocoll

// Set provides a set type.
type Set[T comparable] struct {
	items map[T]bool
}

// NewSet constructs and returns a set of the given type.
func NewSet[T comparable]() Set[T] {
	return Set[T]{make(map[T]bool)}
}

// Contains returns whether the Set contains the indicated item.
// TODO(dburger): test.
func (s *Set[T]) Contains(i T) bool {
	_, ok := s.items[i]
	return ok
}

// Add adds the given item to the Set and returns whether the item was added.
// If the item was already present, false is returned.
func (s *Set[T]) Add(i T) bool {
	_, ok := s.items[i]
	s.items[i] = true
	return !ok
}

// AddAll adds each item of items to the Set.
func (s *Set[T]) AddAll(items ...T) {
	for _, i := range items {
		s.Add(i)
	}
}

// Remove removes the given item from the Set and returns whether the item was removed.
// If it was not present, false is returned.
func (s *Set[T]) Remove(i T) bool {
	_, ok := s.items[i]
	if ok {
		delete(s.items, i)
	}
	return ok
}

// Size returns a count of the items in the Set.
// TODO(dburger): test.
func (s *Set[T]) Size() int {
	return len(s.items)
}
