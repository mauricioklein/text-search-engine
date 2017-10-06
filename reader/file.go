package reader

import (
	"os"
	"strings"
)

// File is the in-memory representation of a OS file
type File struct {
	os.FileInfo
	Path    string
	Content string
}

// NewFile returns a new instance of File
func NewFile(path string, info os.FileInfo, content string) File {
	return File{
		FileInfo: info,
		Path:     path,
		Content:  content,
	}
}

// Words split the file content in every whitespace and
// line break and return a slice of words
func (file File) Words() []string {
	oneLineContent := strings.Replace(file.Content, "\n", " ", -1)
	return strings.Split(oneLineContent, " ")
}
