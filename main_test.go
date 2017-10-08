package main

import (
	"bytes"
	"os"
	"sync"
	"testing"

	"github.com/mauricioklein/text-search-engine/ranking"
	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/mauricioklein/text-search-engine/report"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Args = append(os.Args, "-directory", "./test-utils/3-files")

	// IO buffers
	inBuf := bytes.NewBuffer([]byte{})
	outBuf := bytes.NewBuffer([]byte{})

	// overwrite the standard streams from main
	// to use our mocked one
	inStream = inBuf
	outStream = outBuf

	// user input
	inBuf.WriteString("Lorem ipsum\n")
	inBuf.WriteString(QuitSentence + "\n")

	// dispatch main method in a separated routine
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		main()
	}(&wg)
	wg.Wait()

	actual, _ := outBuf.ReadString('\x00')
	expected := `search> file1.txt: 100.00% match
file3.txt: 100.00% match
file2.txt: 50.00% match
search> `

	assert.Equal(t, expected, actual)
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
