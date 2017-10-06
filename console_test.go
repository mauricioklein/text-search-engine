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

func TestConsoleRead(t *testing.T) {
	c, r, _ := NewTestConsole()

	r.WriteString("Foobar\n")

	actual, _ := c.Read()
	expected := "Foobar"
	assert.Equal(t, expected, actual)
}

func TestConsoleWrite(t *testing.T) {
	c, _, w := NewTestConsole()

	c.Write("Foobar")

	actual, _ := w.ReadString('\x00')
	expected := "Foobar"
	assert.Equal(t, expected, actual)
}

//
// TODO: re-enable this test after the results ordering is done
//
func TestConsoleProcess(t *testing.T) {
	t.Skip()

	c, _, w := NewTestConsole()

	c.process("Cat")

	actual, _ := w.ReadString('\x00')
	expected := "file1.txt: 100.00% match\nfile2.txt: 100.00% match\nfile3.txt: 0.00% match\n"
	assert.Equal(t, expected, actual)
}

//
// TODO: re-enable this test after the results ordering is done
//
func TestConsoleRun(t *testing.T) {
	t.Skip()

	c, r, w := NewTestConsole()

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

func NewTestConsole() (Console, *bytes.Buffer, *bytes.Buffer) {
	files, _ := reader.Disk{}.Read("./test-utils/files/")

	reader := bytes.NewBuffer([]byte{})
	writer := bytes.NewBuffer([]byte{})

	return NewConsole(
		files,
		ranking.LevenshteinRanking{},
		report.SimpleReporter{},
		reader,
		writer,
	), reader, writer
}

func dispatchConsole(c Console, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Run()
}
