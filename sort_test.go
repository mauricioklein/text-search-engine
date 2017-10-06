package main

import (
	"os"
	"sort"
	"testing"

	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/stretchr/testify/assert"
)

func TestSortByRank(t *testing.T) {
	rank25 := RankResult{Rank: 0.25}
	rank50 := RankResult{Rank: 0.50}
	rank75 := RankResult{Rank: 0.75}

	ranks := []RankResult{rank75, rank50, rank25}
	sort.Sort(ByScoreAndName(ranks))

	assert.Exactly(t, ranks, []RankResult{rank25, rank50, rank75})
}

func TestSortByName(t *testing.T) {
	f1, _ := os.Open("./test-utils/files/file1.txt")
	f2, _ := os.Open("./test-utils/files/file2.txt")
	f3, _ := os.Open("./test-utils/files/subdir/file3.txt")

	f1stat, _ := f1.Stat()
	f2stat, _ := f2.Stat()
	f3stat, _ := f3.Stat()

	rank1 := RankResult{File: reader.File{FileInfo: f1stat}, Rank: 0.5}
	rank2 := RankResult{File: reader.File{FileInfo: f2stat}, Rank: 0.5}
	rank3 := RankResult{File: reader.File{FileInfo: f3stat}, Rank: 0.5}

	ranks := []RankResult{rank1, rank2, rank3}
	sort.Sort(ByScoreAndName(ranks))

	// sort by name is reverse
	assert.Exactly(t, ranks, []RankResult{rank3, rank2, rank1})
}

func TestSortByRankAndName(t *testing.T) {
	f1, _ := os.Open("./test-utils/files/file1.txt")
	f2, _ := os.Open("./test-utils/files/file2.txt")
	f3, _ := os.Open("./test-utils/files/subdir/file3.txt")

	f1stat, _ := f1.Stat()
	f2stat, _ := f2.Stat()
	f3stat, _ := f3.Stat()

	rank1 := RankResult{File: reader.File{FileInfo: f1stat}, Rank: 0.25}
	rank2 := RankResult{File: reader.File{FileInfo: f2stat}, Rank: 0.50}
	rank3 := RankResult{File: reader.File{FileInfo: f3stat}, Rank: 0.50}

	ranks := []RankResult{rank3, rank2, rank1}
	sort.Sort(ByScoreAndName(ranks))

	// rank1 first (lowest rank)
	// rank3 before rank2 (same rank, but the sort by name is reverse)
	assert.Exactly(t, ranks, []RankResult{rank1, rank3, rank2})
}
