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

func TestConsoleInputStream(t *testing.T) {
	c, r, _, _ := NewTestConsole()

	r.WriteString("Foobar\n")

	actual, _ := c.Read()
	expected := "Foobar"
	assert.Equal(t, expected, actual)
}

func TestConsoleOutputStream(t *testing.T) {
	c, _, w, _ := NewTestConsole()

	c.Write("Foobar")

	actual, _ := w.ReadString('\x00')
	expected := "Foobar"
	assert.Equal(t, expected, actual)
}

func TestConsoleErrorStream(t *testing.T) {
	c, _, _, e := NewTestConsole()

	c.Error("a generic error")

	actual, _ := e.ReadString('\x00')
	expected := "a generic error"
	assert.Equal(t, expected, actual)
}

func TestConsoleProcess(t *testing.T) {
	c, _, w, _ := NewTestConsole()

	c.process("Cat")

	actual, _ := w.ReadString('\x00')
	expected := "file1.txt: 100.00% match\nfile2.txt: 100.00% match\nfile3.txt: 0.00% match\n"
	assert.Equal(t, expected, actual)
}

func TestConsoleRun(t *testing.T) {
	c, r, w, _ := NewTestConsole()

	// write "user input" data to the read stream
	r.Write([]byte("Cat\n")) // actual search sentence
	r.Write([]byte("\\q\n")) // quit command

	// Wait for the run command to finish (due the quit command above)
	var wg sync.WaitGroup
	wg.Add(1)
	go dispatchConsole(c, &wg)
	wg.Wait()

	// Read response from the write stream
	actual, _ := w.ReadString('\x00')
	expected := "search> file1.txt: 100.00% match\nfile2.txt: 100.00% match\nfile3.txt: 0.00% match\nsearch> "

	assert.Equal(t, expected, actual)
}

func NewTestConsole() (Console, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	files, _ := reader.Disk{}.Read("./test-utils/files/")

	inBuf := bytes.NewBuffer([]byte{})
	outBuf := bytes.NewBuffer([]byte{})
	errBuf := bytes.NewBuffer([]byte{})

	return NewConsole(
		files,
		ranking.LevenshteinRanking{},
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
