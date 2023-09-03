package gocoll

type dequeNode[T any] struct {
	val  T
	prev *dequeNode[T]
	next *dequeNode[T]
}

type Deque[T any] struct {
	head *dequeNode[T]
	tail *dequeNode[T]
	size int
}

func NewDeque[T any]() Deque[T] {
	head := dequeNode[T]{}
	tail := dequeNode[T]{}
	head.next = &tail
	tail.prev = &head
	return Deque[T]{&head, &tail, 0}
}

func (d *Deque[T]) AddFirst(val T) {
	n := dequeNode[T]{val, d.head, d.head.next}
	d.head.next.prev = &n
	d.head.next = &n
	d.size++
}

func (d *Deque[T]) AddLast(val T) {
	n := dequeNode[T]{val, d.tail.prev, d.tail}
	d.tail.prev.next = &n
	d.tail.prev = &n
	d.size++
}

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

func (d *Deque[T]) PeekFirst() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}
	return d.head.next.val, true
}

func (d *Deque[T]) PeekLast() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}
	return d.tail.prev.val, true
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque[T]) Visit(f func(T)) {
	n := d.head.next
	for n != d.tail {
		f(n.val)
		n = n.next
	}
}
