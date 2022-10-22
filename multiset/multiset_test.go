package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(1)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 0, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	ms.Add(2)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 1, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	ms.Add(2)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 0, ms.Count(3))
	ms.Add(3)
	assert.Equal(t, 1, ms.Count(1))
	assert.Equal(t, 2, ms.Count(2))
	assert.Equal(t, 1, ms.Count(3))
	assert.Equal(t, 0, ms.Count(4))
}

func TestRemove(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(1)
	ms.Add(2)
	ms.Add(2)
	ms.Add(3)
}
