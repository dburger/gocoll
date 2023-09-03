package gocoll

import "fmt"

type linkedListNode[T comparable] struct {
	val  T
	next *linkedListNode[T]
}

type linkedList[T comparable] struct {
	head *linkedListNode[T]
}

func NewLinkedList[T comparable]() *linkedList[T] {
	return &linkedList[T]{nil}
}

func (ll linkedList[T]) Display() {
	for n := ll.head; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}

func (ll *linkedList[T]) Add(val T) {
	n := &linkedListNode[T]{val, ll.head}
	ll.head = n
}

func (ll *linkedList[T]) Remove(val T) {
	var prev *linkedListNode[T]
	curr := ll.head
	for curr != nil && curr.val != val {
		prev = curr
		curr = curr.next
	}
	if curr != nil {
		if prev == nil {
			ll.head = curr.next
		} else {
			prev.next = curr.next
		}
	}
}

/*
type linkedList[T any] struct {
	head *linkedListNode[T]
	comp func(a, b T) bool
}

func NewLinkedList[T any](comp func(a, b T) bool) *linkedList[T] {
	return &linkedList[T]{nil, comp}
}

func (ll linkedList[T]) Display() {
	for n := ll.head; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}

func (ll *linkedList[T]) Add(val T) {
	n := &linkedListNode[T]{val, ll.head}
	ll.head = n
}

func (ll *linkedList[T]) Remove(val T) {
	var prev *linkedListNode[T]
	curr := ll.head
	for curr != nil && !ll.comp(curr.val, val) {
		prev = curr
		curr = curr.next
	}
	if curr != nil {
		if prev == nil {
			ll.head = curr.next
		} else {
			prev.next = curr.next
		}
	}
}
*/
