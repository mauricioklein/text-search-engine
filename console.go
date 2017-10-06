package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/mauricioklein/text-search-engine/ranking"
	"github.com/mauricioklein/text-search-engine/report"
)

// QuitSentence defines the sentence, read from the
// input stream, to quit the console
const QuitSentence = "\\q"

// Console defines an instance of the
// interactive console
type Console struct {
	Processor    ranking.Processor
	Reporter     report.Reporter
	InputStream  *bufio.Reader
	OutputStream *bufio.Writer
	ErrorStream  *bufio.Writer
}

// NewConsole creates a new instance of Console
func NewConsole(processor ranking.Processor, reporter report.Reporter, inputStream io.Reader, outputStream io.Writer, errStream io.Writer) Console {
	inputBuffer := bufio.NewReader(inputStream)
	outputBuffer := bufio.NewWriter(outputStream)
	errBuffer := bufio.NewWriter(errStream)

	return Console{
		Processor:    processor,
		Reporter:     reporter,
		InputStream:  inputBuffer,
		OutputStream: outputBuffer,
		ErrorStream:  errBuffer,
	}
}

// Write writes a string to the console's output stream
func (c Console) Write(line string) {
	c.OutputStream.Write([]byte(line))
	c.OutputStream.Flush()
}

func (c Console) Error(line string) {
	c.ErrorStream.Write([]byte(line))
	c.ErrorStream.Flush()
}

// Read reads a string from the console's input stream
func (c Console) Read() (string, error) {
	rawInput, err := c.InputStream.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Replace(rawInput, "\n", "", -1), nil
}

// Flush flushes the content of output stream
func (c Console) Flush() {
	c.OutputStream.Flush()
	c.ErrorStream.Flush()
}

// Run executes and controls the IO of
// the console with the user
func (c Console) Run() {
	for {
		c.Write("search> ")

		userInput, err := c.Read()
		if err != nil {
			c.Error(fmt.Sprintf("Failed to read input: %s", err))
			continue
		}

		if isStopCondition(userInput) {
			break
		}

		// calculate the ranks
		ranks := c.Processor.Calculate(userInput)

		// order ranks by score/filename
		sort.Sort(
			sort.Reverse(
				ranking.ByScoreAndName(ranks),
			),
		)

		// print out the results
		for _, rank := range ranks {
			c.ReportRank(rank)
		}

		// flush the output stream
		c.Flush()
	}
}

// ReportRank reports a given result
func (c Console) ReportRank(rr ranking.RankResult) {
	c.Reporter.Report(
		c.OutputStream,
		rr.File.Name(),
		rr.Score,
	)
}

// isStopCondition checks if the input stream contains
// the console's stop sentence
func isStopCondition(userInput string) bool {
	return userInput == QuitSentence
}
