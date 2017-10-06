package main

import (
	"bytes"
	"sync"
	"testing"

	"github.com/mauricioklein/text-search-engine/ranking"
	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/mauricioklein/text-search-engine/report"
	"github.com/stretchr/testify/assert"
)

func TestConsoleInputStreamSuccess(t *testing.T) {
	c, r, _, _ := NewTestConsole("./test-utils/3-files/")

	r.WriteString("Foobar\n")

	actual, _ := c.Read()
	expected := "Foobar"
	assert.Equal(t, expected, actual)
}

func TestConsoleInputStreamError(t *testing.T) {
	c, r, _, _ := NewTestConsole("./test-utils/3-files/")

	r.WriteString("Foobar")

	_, err := c.Read()
	assert.Error(t, err)
}

func TestConsoleOutputStream(t *testing.T) {
	c, _, w, _ := NewTestConsole("./test-utils/3-files/")

	c.Write("Foobar")

	actual, _ := w.ReadString('\x00')
	expected := "Foobar"
	assert.Equal(t, expected, actual)
}

func TestConsoleErrorStream(t *testing.T) {
	c, _, _, e := NewTestConsole("./test-utils/3-files/")

	c.Error("a generic error")

	actual, _ := e.ReadString('\x00')
	expected := "a generic error"
	assert.Equal(t, expected, actual)
}

func TestConsoleRun3Files(t *testing.T) {
	c, r, w, _ := NewTestConsole("./test-utils/3-files/")

	// write "user input" data to the read stream
	r.Write([]byte("Ca\n"))  // actual search sentence
	r.Write([]byte("\\q\n")) // quit command

	// Wait for the run command to finish (due the quit command above)
	var wg sync.WaitGroup
	wg.Add(1)
	go dispatchConsole(c, &wg)
	wg.Wait()

	// Read response from the write stream
	actual, _ := w.ReadString('\x00')
	expected := `search> file1.txt: 66.67% match
file2.txt: 66.67% match
file3.txt: 0.00% match
search> `

	assert.Equal(t, expected, actual)
}

func TestConsoleRun11Files(t *testing.T) {
	c, r, w, _ := NewTestConsole("./test-utils/11-files/")

	// write "user input" data to the read stream
	r.Write([]byte("Ca\n"))  // actual search sentence
	r.Write([]byte("\\q\n")) // quit command

	// Wait for the run command to finish (due the quit command above)
	var wg sync.WaitGroup
	wg.Add(1)
	go dispatchConsole(c, &wg)
	wg.Wait()

	// Read response from the write stream
	actual, _ := w.ReadString('\x00')

	// file11.txt should not be present, since the result
	// should display only the top 10
	expected := `search> file1.txt: 66.67% match
file2.txt: 66.67% match
file10.txt: 0.00% match
file3.txt: 0.00% match
file4.txt: 0.00% match
file5.txt: 0.00% match
file6.txt: 0.00% match
file7.txt: 0.00% match
file8.txt: 0.00% match
file9.txt: 0.00% match
search> `

	// should filter the result to the top 10
	assert.Equal(t, expected, actual)
}

func NewTestConsole(dirPath string) (Console, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	files, _ := reader.Disk{}.Read(dirPath)

	nWorkers := 3
	processor := ranking.NewProcessor(files, nWorkers, ranking.LevenshteinRanking{})
	inBuf := bytes.NewBuffer([]byte{})
	outBuf := bytes.NewBuffer([]byte{})
	errBuf := bytes.NewBuffer([]byte{})

	return NewConsole(
		processor,
		report.SimpleReporter{},
		inBuf,
		outBuf,
		errBuf,
	), inBuf, outBuf, errBuf
}

func dispatchConsole(c Console, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Run()
}
