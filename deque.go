package gocoll

type dequeNode[T any] struct {
	val  T
	prev *dequeNode[T]
	next *dequeNode[T]
}

// Deque is a double ended queue data structure.
type Deque[T any] struct {
	head *dequeNode[T]
	tail *dequeNode[T]
	size int
}

// NewDeque constructs a new Deque of the given type.
func NewDeque[T any]() Deque[T] {
	head := dequeNode[T]{}
	tail := dequeNode[T]{}
	head.next = &tail
	tail.prev = &head
	return Deque[T]{&head, &tail, 0}
}

// AddFirst adds an element to the head of the Deque.
func (d *Deque[T]) AddFirst(val T) {
	n := dequeNode[T]{val, d.head, d.head.next}
	d.head.next.prev = &n
	d.head.next = &n
	d.size++
}

// AddLast adds an element to the tail of the Deque.
func (d *Deque[T]) AddLast(val T) {
	n := dequeNode[T]{val, d.tail.prev, d.tail}
	d.tail.prev.next = &n
	d.tail.prev = &n
	d.size++
}

// RemoveFirst removes and returns the head of the Deque.
// If the Deque is empty, a zero element and false is returned.
func (d *Deque[T]) RemoveFirst() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}
	val := d.head.next.val
	d.head.next = d.head.next.next
	d.head.next.prev = d.head
	d.size--
	return val, true
}

// RemoveLast removes and returns the tail of the Deque.
// If the Deque is empty, a zero element and false is returned.
func (d *Deque[T]) RemoveLast() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}
	val := d.tail.prev.val
	d.tail.prev = d.tail.prev.prev
	d.tail.prev.next = d.tail
	d.size--
	return val, true
}

// PeekFirst returns the head of the Deque.
// If the Deque is empty, a zero element and false is returned.
func (d *Deque[T]) PeekFirst() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}
	return d.head.next.val, true
}

// PeekLast returns the head of the Deque.
// If the Deque is empty, a zero element and false is returned.
func (d *Deque[T]) PeekLast() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}
	return d.tail.prev.val, true
}

// Size returns the number of elements in the Deque.
func (d *Deque[T]) Size() int {
	return d.size
}

// IsEmpty returns whether the Deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Visit executes func f on every element of the Deque.
func (d *Deque[T]) Visit(f func(T)) {
	n := d.head.next
	for n != d.tail {
		f(n.val)
		n = n.next
	}
}
