package main

import (
	"bytes"
	"testing"

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

func TestConsoleProcess(t *testing.T) {
	c, _, w := NewTestConsole()

	c.process("Cat")

	actual, _ := w.ReadString('\x00')
	expected := "file1.txt: 100.000000\nfile2.txt: 100.000000\n"
	assert.Equal(t, expected, actual)
}

func NewTestConsole() (Console, *bytes.Buffer, *bytes.Buffer) {
	files, _ := DiskReader{}.Read("./test-utils/files/")

	reader := bytes.NewBuffer([]byte{})
	writer := bytes.NewBuffer([]byte{})

	return NewConsole(
		files,
		LevenshteinRanking{},
		reader,
		writer,
	), reader, writer
}
