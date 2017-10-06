package reader

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Disk defines a implementation of the
// Reader interface to read files from Disk
type Disk struct{}

// Read reads the files from disk and translate it
// to the internal File structure
func (d Disk) Read(path string) ([]File, error) {
	var files []File
	err := processDirectory(path, &files)
	return files, err
}

// processDirectory processes the directory refered by
// "path" and, recursively, append the files do "acc"
func processDirectory(path string, acc *[]File) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			processDirectory(buildPath(path, file.Name()), acc)
		} else {
			processFile(path, file, acc)
		}
	}

	return nil
}

// processFile loads the file content and return an instance of File
func processFile(basePath string, file os.FileInfo, acc *[]File) error {
	content, err := ioutil.ReadFile(buildPath(basePath, file.Name()))
	if err != nil {
		return err
	}

	*acc = append(*acc, File{
		FileInfo: file,
		Path:     basePath,
		Content:  string(content),
	})

	return nil
}

// buildPath returns the path composed by the base path
// and the file path
func buildPath(basePath, filePath string) string {
	return fmt.Sprintf("%s/%s", basePath, filePath)
}
