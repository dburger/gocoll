package gocoll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDeque() Deque[int] {
	d := NewDeque[int]()
	d.AddFirst(1)
	d.AddLast(2)
	d.AddLast(3)
	d.AddLast(4)
	d.AddLast(5)
	return d
}

func TestPeek(t *testing.T) {
	d := makeDeque()
	val, ok := d.PeekFirst()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	val, ok = d.PeekLast()
	assert.True(t, ok)
	assert.Equal(t, 5, val)
}

func TestRemove(t *testing.T) {
	d := makeDeque()
	val, ok := d.RemoveFirst()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 4, d.Size())
	val, ok = d.RemoveLast()
	assert.True(t, ok)
	assert.Equal(t, 5, val)
	assert.Equal(t, 3, d.Size())
}

func TestIsEmpty(t *testing.T) {
	d := makeDeque()
	assert.False(t, d.IsEmpty())
	d.RemoveFirst()
	d.RemoveLast()
	assert.Equal(t, 3, d.Size())
	d.RemoveFirst()
	d.RemoveLast()
	d.RemoveFirst()
	assert.True(t, d.IsEmpty())
}
