package main

import (
	"fmt"
	"io/ioutil"
)

// Reader defines an interface that implements the
// methods to read files from some resource and translate
// it to File instances
type Reader interface {
	Read(path string) ([]File, error)
}

// DiskReader defines a implementation of the
// Reader interface to read files from Disk
type DiskReader struct{}

// Read reads the files from disk and translate it
// to the internal File structure
func (dr DiskReader) Read(path string) ([]File, error) {
	var files []File

	memFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, memFile := range memFiles {
		if memFile.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, memFile.Name()))
		if err != nil {
			return nil, err
		}

		files = append(files, File{
			FileInfo: memFile,
			Path:     path,
			Content:  string(content),
		})
	}

	return files, nil
}
