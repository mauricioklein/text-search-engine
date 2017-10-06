package reader

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Disk defines a implementation of the
// Reader interface to read files from Disk
type Disk struct{}

// Read reads the files from disk and translate it
// to the internal File structure
func (d Disk) Read(path string) ([]File, error) {
	var files []File
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() {
			return nil
		}

		return processFile(path, info, &files)
	})
	return files, err
}

// processFile loads the file content and return an instance of File
func processFile(filepath string, file os.FileInfo, acc *[]File) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	*acc = append(*acc, NewFile(
		path.Dir(filepath),
		file,
		string(content),
	))

	return nil
}
