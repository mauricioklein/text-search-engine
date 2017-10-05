package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevenshteinRanking(t *testing.T) {
	rank := LevenshteinRanking{}
	text := "The quick brown fox jumps over the lazy dog"
	sentence := ""

	// Case 1: 1 exact word match
	sentence = "fox"
	assert.Equal(t, 1.0, rank.Calculate(text, sentence))

	// Case 2: 2 exact word match
	sentence = "brown fox"
	assert.Equal(t, 1.0, rank.Calculate(text, sentence))

	// Case 3: 1 partial word match
	sentence = "brw" // rank = 1 - (2/5) = 60% match
	assert.Equal(t, 0.6, rank.Calculate(text, sentence))

	// Case 4: 2 partial word match
	sentence = "brw fx" // Changes = (1 - (2/5)) + (1 - (1/3)) = 63.333% match
	assert.InDelta(t, 0.633, rank.Calculate(text, sentence), 0.001)

	// Case 5: No match
	sentence = "xyz"
	assert.Equal(t, 0.0, rank.Calculate(text, sentence))
}
