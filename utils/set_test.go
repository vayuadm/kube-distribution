package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_NotContainsInEmptySet(t *testing.T) {

	assert.False(t, NewSet().Contains("bla"))
}

func TestSet_Contains(t *testing.T) {

	s := NewSet()
	s.Add("bla")
	assert.True(t, s.Contains("bla"))
}

func TestSet_NotContains(t *testing.T) {

	s := NewSet()
	s.Add("bla")
	assert.False(t, s.Contains("blabla"))
}

func TestSet_ToArray(t *testing.T) {

	const x1 = "hello"

	s := NewSet()
	s.Add(x1)
	arr := s.ToArray()

	assert.Equal(t, x1, arr[0])
}
