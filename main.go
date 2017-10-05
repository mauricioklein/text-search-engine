package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var reader Reader = DiskReader{}
	var reporter Reporter = SimpleReporter{}
	var rank Ranking = LevenshteinRanking{}

	var path = "./test-utils/files/"

	files, err := reader.Read(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d file(s) read in the directory %s\n", len(files), path)

	NewConsole(
		files,
		rank,
		os.Stdin,
		os.Stderr,
		reporter,
	).Run()
}
