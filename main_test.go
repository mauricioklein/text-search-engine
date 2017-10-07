package main

import (
	"os"
	"testing"

	"github.com/mauricioklein/text-search-engine/ranking"
	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/mauricioklein/text-search-engine/report"
	"github.com/stretchr/testify/assert"
)

func TestParseCliFlags(t *testing.T) {
	os.Args = append(os.Args,
		"-directory", "./test-utils/3-files",
		"-reader", "usb",
		"-reporter", "my_custom_reporter",
		"-rank", "foobar_rank",
		"-workers", "5",
	)

	subject := parseCliFlags()

	assert.Equal(t, subject.DirPath, "./test-utils/3-files")
	assert.Equal(t, subject.Reader, "usb")
	assert.Equal(t, subject.Reporter, "my_custom_reporter")
	assert.Equal(t, subject.RankAlgo, "foobar_rank")
	assert.Equal(t, subject.NWorkers, 5)
}

func TestInstantiateReader(t *testing.T) {
	assert.IsType(t, reader.Disk{}, instantiateReader(""))
}

func TestInstantiateReporter(t *testing.T) {
	assert.IsType(t, report.SimpleReporter{}, instantiateReporter(""))
}

func TestInstantiateRankingAlgorithm(t *testing.T) {
	assert.IsType(t, ranking.LevenshteinRanking{}, instantiateRankingAlgorithm(""))
}
