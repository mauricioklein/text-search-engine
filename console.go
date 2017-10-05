package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
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

// RankResult defines the result of a
// rank canculation for a specific file
type RankResult struct {
	File File
	Rank float64
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
	for {
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
	jobs := make(chan File, len(c.Files))
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

	// write out the results
	close(results)
	for r := range results {
		c.Write(fmt.Sprintf("%s: %f\n", r.File.Name(), r.Rank*100.0))
	}
}

func (c Console) worker(jobs <-chan File, results chan<- RankResult, sentence string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		results <- RankResult{
			File: file,
			Rank: c.Rank.Calculate(file.Content, sentence),
		}
	}
}

func isStopCondition(userInput string) bool {
	return userInput == QuitSentence
}
