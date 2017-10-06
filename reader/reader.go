package reader

import "os"

// File is the in-memory representation of a OS file
type File struct {
	os.FileInfo
	Path    string
	Content string
}

// Reader defines an interface that implements the
// methods to read files from some resource and translate
// it to File instances
type Reader interface {
	Read(path string) ([]File, error)
}
