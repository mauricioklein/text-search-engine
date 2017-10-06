package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mauricioklein/text-search-engine/ranking"
	"github.com/mauricioklein/text-search-engine/reader"
	"github.com/mauricioklein/text-search-engine/report"
)

// CliArgs defines the command line arguments
// provided in the program execution
type CliArgs struct {
	DirPath  string
	Reader   string
	Reporter string
	RankAlgo string
	NWorkers int
}

func main() {
	// parse command line arguments
	args := parseCliFlags()

	// instantiate the necessary resources
	reader := instantiateReader(args.Reader)
	reporter := instantiateReporter(args.Reporter)
	rankAlgo := instantiateRankingAlgorithm(args.RankAlgo)
	nWorkers := args.NWorkers

	files, err := reader.Read(args.DirPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d file(s) read in the directory %s\n", len(files), args.DirPath)

	NewConsole(
		files,
		rankAlgo,
		reporter,
		nWorkers,
		os.Stdin,
		os.Stdout,
		os.Stderr,
	).Run()
}

func parseCliFlags() CliArgs {
	path := flag.String("directory", "", "Directory to be read")
	reader := flag.String("reader", "disk", "The file reader to be used")
	reporter := flag.String("reporter", "simple", "The result reporter to be used")
	rankAlgo := flag.String("rank", "levenshtein", "The rank algorithm to be used")
	nWorkers := flag.Int("workers", 3, "Number of paralel workers")
	flag.Parse()

	return CliArgs{
		DirPath:  *path,
		Reader:   *reader,
		Reporter: *reporter,
		RankAlgo: *rankAlgo,
		NWorkers: *nWorkers,
	}
}

func instantiateReader(readerType string) reader.Reader {
	switch readerType {
	default:
		return reader.Disk{}
	}
}

func instantiateReporter(reporterType string) report.Reporter {
	switch reporterType {
	default:
		return report.SimpleReporter{}
	}
}

func instantiateRankingAlgorithm(algType string) ranking.Algorithm {
	switch algType {
	default:
		return ranking.LevenshteinRanking{}
	}
}
