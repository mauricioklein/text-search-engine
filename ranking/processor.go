package ranking

import (
	"sync"

	"github.com/mauricioklein/text-search-engine/reader"
)

// Processor defines a instance of the processor
type Processor struct {
	Files         []reader.File
	NWorkers      int
	RankAlgorithm Algorithm
}

// RankResult defines the result of a
// rank canculation for a specific file
type RankResult struct {
	File reader.File
	Rank float64
}

// NewProcessor creates a new processor instance
func NewProcessor(files []reader.File, nWorkers int, rankAlg Algorithm) Processor {
	return Processor{
		Files:         files,
		NWorkers:      nWorkers,
		RankAlgorithm: rankAlg,
	}
}

// Calculate calculates the rank for all the files hold by
// processor in comparison with a given sentence
func (p Processor) Calculate(sentence string) []RankResult {
	jobs := make(chan reader.File, len(p.Files))
	results := make(chan RankResult, len(p.Files))

	// create workers
	var wg sync.WaitGroup
	for i := 0; i < p.NWorkers; i++ {
		wg.Add(1)
		go p.worker(jobs, results, sentence, &wg)
	}

	// add jobs to the channel
	for _, file := range p.Files {
		jobs <- file
	}
	close(jobs)

	// wait for all workers to finish
	wg.Wait()

	// process the results
	close(results)

	return toSlice(results)
}

// worker performs the rank calculation for a given file
func (p Processor) worker(jobs <-chan reader.File, results chan<- RankResult, sentence string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		results <- RankResult{
			File: file,
			Rank: p.RankAlgorithm.Calculate(file.Content, sentence),
		}
	}
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
