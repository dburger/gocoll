package collections

type set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() set[T] {
	return set[T]{make(map[T]bool)}
}

func (s *set[T]) contains(i T) bool {
	_, ok := s.items[i]
	return ok
}

func (s *set[T]) add(i T) bool {
	_, ok := s.items[i]
	return ok
}

func (s *set[T]) remove(i T) bool {
	_, ok := s.items[i]
	if ok {
		delete(s.items, i)
	}
	return ok
}
