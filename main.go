package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var path = "./test-utils/files/"
	var reader Reader = DiskReader{}

	files, err := reader.Read(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d file(s) read in the directory %s\n", len(files), path)

	NewConsole(
		files,
		LevenshteinRanking{},
		os.Stdin,
		os.Stderr,
	).Run()
}
