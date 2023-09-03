package gocoll

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{make(map[T]bool)}
}

// TODO(dburger): test.
func (s *Set[T]) Contains(i T) bool {
	_, ok := s.items[i]
	return ok
}

func (s *Set[T]) Add(i T) bool {
	_, ok := s.items[i]
	s.items[i] = true
	return ok
}

func (s *Set[T]) Remove(i T) bool {
	_, ok := s.items[i]
	if ok {
		delete(s.items, i)
	}
	return ok
}

// TODO(dburger): test.
func (s *Set[T]) Size() int {
	return len(s.items)
}
