package ranking

import (
	"testing"

	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	files, _ := reader.Disk{}.Read("../test-utils/3-files/")
	processor := NewProcessor(files, 3, LevenshteinRanking{})

	ranks := processor.Calculate("Cat")

	// Rank: file1.txt
	rank := findByName(ranks, "file1.txt")
	assert.Equal(t, rank.File.Name(), "file1.txt")
	assert.InDelta(t, 0.333, rank.Score, 0.001)

	// Rank: file2.txt
	rank = findByName(ranks, "file2.txt")
	assert.Equal(t, rank.File.Name(), "file2.txt")
	assert.InDelta(t, 0.333, rank.Score, 0.001)

	// Rank: file3.txt
	rank = findByName(ranks, "file3.txt")
	assert.Equal(t, rank.File.Name(), "file3.txt")
	assert.InDelta(t, 0.333, rank.Score, 0.001)
}

func findByName(ranks []RankResult, filename string) *RankResult {
	for _, rank := range ranks {
		if rank.File.Name() == filename {
			return &rank
		}
	}

	return nil
}
