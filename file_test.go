package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	file := NewFile("foo/bar/", nil, "FooBar")

	assert.Equal(t, "foo/bar/", file.Path)
	assert.Equal(t, "FooBar", file.Content)
}

func TestWords(t *testing.T) {
	words := File{Content: "Word1 Word2\nWord3"}.Words()

	assert.Exactly(t, []string{"Word1", "Word2", "Word3"}, words)
}
