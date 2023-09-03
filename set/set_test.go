package gocoll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	s.Add(2)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	s.Add(2)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))

	s.Add(3)
	assert.Equal(t, 3, s.Size())
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}

func TestRemove(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Size())
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Equal(t, 3, s.Size())

	s.Remove(1)
	assert.Equal(t, 2, s.Size())
	assert.False(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	s.Remove(2)
	assert.Equal(t, 1, s.Size())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.True(t, s.Contains(3))

	s.Remove(3)
	assert.Equal(t, 0, s.Size())
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains(3))
}
