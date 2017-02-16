package main

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

func TestSetNotContains(t *testing.T) {

	s := NewSet()
	s.Add("bla")
	assert.False(t, s.Contains("blabla"))
}