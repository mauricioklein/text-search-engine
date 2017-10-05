package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// QuitSentence defines the sentence, read from the
// input stream, to quit the console
const QuitSentence = "\\q"

// Console defines an instance of the
// interactive console
type Console struct {
	Files        []File
	Rank         Ranking
	InputStream  *bufio.Reader
	OutputStream *bufio.Writer
}

// NewConsole creates a new instance of Console
func NewConsole(files []File, rank Ranking, inputStream io.Reader, outputStream io.Writer) Console {
	input := bufio.NewReader(inputStream)
	output := bufio.NewWriter(outputStream)

	return Console{
		Files:        files,
		Rank:         rank,
		InputStream:  input,
		OutputStream: output,
	}
}

// Write writes a string to the console's output stream
func (c Console) Write(line string) {
	c.OutputStream.Write([]byte(line))
	c.OutputStream.Flush()
}

// Read reads a string from the console's input stream
func (c Console) Read() (string, error) {
	rawInput, err := c.InputStream.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Replace(rawInput, "\n", "", -1), nil
}

// Run executes and controls the IO of
// the console with the user
func (c Console) Run() {
	for true {
		c.Write("search> ")

		userInput, err := c.Read()
		if err != nil {
			fmt.Printf("Failed to read input: %s", err)
			continue
		}

		if isStopCondition(userInput) {
			break
		}

		c.process(userInput)
	}
}

func (c Console) process(sentence string) {
	for _, file := range c.Files {
		rank := c.Rank.Calculate(file.Content, sentence) * 100.0
		c.Write(fmt.Sprintf("%s: %f\n", file.Name(), rank))
	}
}

func isStopCondition(userInput string) bool {
	return userInput == QuitSentence
}