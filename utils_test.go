package gocoll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatFunc(t *testing.T) {
	expected := []int{42, 42, 42, 42, 42, 42}
	fortyTwos := RepeatFunc(func() int { return 42 }, 6)

	assert.Equal(t, expected, fortyTwos)
}

func TestRepeatValues(t *testing.T) {
	expected := []int{6, 6, 6}
	sixes := RepeatValues(6, 3)

	assert.Equal(t, expected, sixes)
}
