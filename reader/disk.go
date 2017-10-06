package reader

import (
	"fmt"
	"io/ioutil"
)

// Disk defines a implementation of the
// Reader interface to read files from Disk
type Disk struct{}

// Read reads the files from disk and translate it
// to the internal File structure
func (d Disk) Read(path string) ([]File, error) {
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
