package reader

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromDiskSuccess(t *testing.T) {
	files, err := Disk{}.Read("../test-utils/files/")
	if err != nil {
		t.Error(err)
	}

	assert.Len(t, files, 2)

	// file1.txt
	file := files[0]
	assert.Implements(t, (*os.FileInfo)(nil), file.FileInfo)
	assert.Equal(t, "../test-utils/files/", file.Path)
	assert.Equal(t, fileContent("../test-utils/files/file1.txt"), file.Content)

	// file1.txt
	file = files[1]
	assert.Implements(t, (*os.FileInfo)(nil), file.FileInfo)
	assert.Equal(t, "../test-utils/files/", file.Path)
	assert.Equal(t, fileContent("../test-utils/files/file2.txt"), file.Content)
}

func TestLoadFromDiskFailure(t *testing.T) {
	_, err := Disk{}.Read("./foo/bar")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func fileContent(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return string(bytes)
}
