package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/mauricioklein/text-search-engine/ranking"
	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/mauricioklein/text-search-engine/report"
)

// QuitSentence defines the sentence, read from the
// input stream, to quit the console
const QuitSentence = "\\q"

// Console defines an instance of the
// interactive console
type Console struct {
	Files        []reader.File
	Algorithm    ranking.Algorithm
	InputStream  *bufio.Reader
	OutputStream *bufio.Writer
	ErrorStream  *bufio.Writer
	Reporter     report.Reporter
}

// RankResult defines the result of a
// rank canculation for a specific file
type RankResult struct {
	File reader.File
	Rank float64
}

// NewConsole creates a new instance of Console
func NewConsole(files []reader.File, algo ranking.Algorithm, reporter report.Reporter, inputStream io.Reader, outputStream io.Writer, errStream io.Writer) Console {
	inputBuffer := bufio.NewReader(inputStream)
	outputBuffer := bufio.NewWriter(outputStream)
	errBuffer := bufio.NewWriter(errStream)

	return Console{
		Files:        files,
		Algorithm:    algo,
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

		c.process(userInput)
	}
}

func (c Console) process(sentence string) {
	jobs := make(chan reader.File, len(c.Files))
	results := make(chan RankResult, len(c.Files))

	// create workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go c.worker(jobs, results, sentence, &wg)
	}

	// add jobs to the channel
	for _, file := range c.Files {
		jobs <- file
	}
	close(jobs)

	// wait for all workers to finish
	wg.Wait()

	// process the results
	close(results)
	ranks := toSlice(results)

	// sort the results by rank/filename
	sort.Sort(
		sort.Reverse(
			ByScoreAndName(ranks),
		),
	)

	// final report
	for _, r := range ranks {
		c.ReportRank(r)
	}

	// Flush the output stream, otherwise the
	// results aren't printed on console
	c.Flush()
}

// ReportRank reports a given result
func (c Console) ReportRank(rr RankResult) {
	c.Reporter.Report(
		c.OutputStream,
		rr.File.Name(),
		rr.Rank,
	)
}

// worker performs the rank calculation for a given file
func (c Console) worker(jobs <-chan reader.File, results chan<- RankResult, sentence string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		results <- RankResult{
			File: file,
			Rank: c.Algorithm.Calculate(file.Content, sentence),
		}
	}
}

// isStopCondition checks if the input stream contains
// the console's stop sentence
func isStopCondition(userInput string) bool {
	return userInput == QuitSentence
}

// toSlice returns a slice with all the elements
// in the channel
func toSlice(c <-chan RankResult) []RankResult {
	var s []RankResult

	for r := range c {
		s = append(s, r)
	}

	return s
}
