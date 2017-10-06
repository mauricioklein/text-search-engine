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
	Rank     string
}

func main() {
	// parse command line arguments
	args := parseCliFlags()

	// instantiate the necessary resources
	reader := instantiateReader(args.Reader)
	reporter := instantiateReporter(args.Reporter)
	rank := instantiateRankingAlgorithm(args.Rank)

	files, err := reader.Read(args.DirPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d file(s) read in the directory %s\n", len(files), args.DirPath)

	NewConsole(
		files,
		rank,
		reporter,
		os.Stdin,
		os.Stdout,
		os.Stderr,
	).Run()
}

func parseCliFlags() CliArgs {
	path := flag.String("directory", "", "Directory to be read")
	reader := flag.String("reader", "disk", "The file reader to be used")
	reporter := flag.String("reporter", "simple", "The result reporter to be used")
	rank := flag.String("rank", "levenshtein", "The rank algorithm to be used")
	flag.Parse()

	return CliArgs{
		DirPath:  *path,
		Reader:   *reader,
		Reporter: *reporter,
		Rank:     *rank,
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
